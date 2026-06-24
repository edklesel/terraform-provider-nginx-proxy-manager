# Copyright IBM Corp. 2021, 2026
# SPDX-License-Identifier: MPL-2.0

import os, requests, re
from collections import defaultdict
from ruamel.yaml import YAML

OWNER_REPO = "NginxProxyManager/nginx-proxy-manager"
WORKFLOW_PATH = ".github/workflows/test.yml"
JOB_NAME = "test-sdk"
MIN_VERSION = (2, 15, 0)  # Always included as a floor
README_PATH = "README.md"

# ── Fetch releases ────────────────────────────────────────────────
token = os.environ["GITHUB_TOKEN"]
response = requests.get(
    f"https://api.github.com/repos/{OWNER_REPO}/releases",
    headers={
        "Authorization": f"Bearer {token}",
        "Accept": "application/vnd.github+json"
    },
    params={"per_page": 100}
)
response.raise_for_status()

# ── Parse versions, skip pre-releases and drafts ──────────────────
versions = []
for r in response.json():
    if r["prerelease"] or r["draft"]:
        continue
    tag = r["tag_name"].lstrip("v")
    parts = tag.split(".")
    if len(parts) == 3:
        versions.append(tuple(int(x) for x in parts))

versions.sort(reverse=True)

# ── Find the 5 most recent distinct minors ────────────────────────
seen_minors = []
by_minor = defaultdict(list)
for v in versions:
    minor_key = (v[0], v[1])
    if minor_key not in seen_minors:
        seen_minors.append(minor_key)
    by_minor[minor_key].append(v)
    if len(seen_minors) == 5:
        break

latest_minor = seen_minors[0]
older_minors = seen_minors[1:]

# ── Build version list ────────────────────────────────────────────
result = []

# Older minors: bare "maj.min", filtered below MIN_VERSION minor
for m in reversed(older_minors):
    if m >= (MIN_VERSION[0], MIN_VERSION[1]):
        result.append(f"{m[0]}.{m[1]}")

# Latest minor: last 5 patches, oldest first, filtered below MIN_VERSION
latest_patches = [v for v in sorted(by_minor[latest_minor])[-5:] if v >= MIN_VERSION]
for v in latest_patches:
    result.append(f"{v[0]}.{v[1]}.{v[2]}")

# Force MIN_VERSION into the list if not already present
min_str = f"{MIN_VERSION[0]}.{MIN_VERSION[1]}.{MIN_VERSION[2]}"
min_minor_str = f"{MIN_VERSION[0]}.{MIN_VERSION[1]}"
if min_str not in result and min_minor_str not in result:
    result.insert(0, min_str)

# ── Load, update, and write the workflow file ─────────────────────
yaml = YAML()
yaml.preserve_quotes = True

with open(WORKFLOW_PATH) as f:
    workflow = yaml.load(f)

current = [str(v) for v in workflow["jobs"][JOB_NAME]["strategy"]["matrix"]["NPM_VERSION"]]

if current == result:
    print("Versions unchanged, skipping update.")
else:
    workflow["jobs"][JOB_NAME]["strategy"]["matrix"]["NPM_VERSION"] = result

    with open(WORKFLOW_PATH, "w") as f:
        yaml.dump(workflow, f)

    print("Updated VERSION matrix:")
    for v in result:
        print(f"  - {v}")

# ── Update README ─────────────────────────────────────────────────
with open(README_PATH) as f:
    readme = f.read()

new_list = "\n".join(f"- {v}" for v in result)
new_section = f"The provider is tested against the following versions in CI:\n{new_list}"

updated_readme = re.sub(
    r"The provider is tested against the following versions in CI:\n(?:- .+\n?)+",
    new_section + "\n",
    readme
)

if updated_readme == readme:
    print("README versions unchanged, skipping.")
else:
    with open(README_PATH, "w") as f:
        f.write(updated_readme)
    print("Updated README versions.")
