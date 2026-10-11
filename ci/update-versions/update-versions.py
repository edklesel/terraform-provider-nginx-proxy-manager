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

# ── Fetch releases (all pages) ────────────────────────────────────
token = os.environ["GITHUB_TOKEN"]
headers = {
    "Authorization": f"Bearer {token}",
    "Accept": "application/vnd.github+json",
}

releases = []
url = f"https://api.github.com/repos/{OWNER_REPO}/releases"
params = {"per_page": 100}
while url:
    response = requests.get(url, headers=headers, params=params)
    response.raise_for_status()
    releases.extend(response.json())
    url = response.links.get("next", {}).get("url")
    params = None  # the "next" URL already carries the query string

# ── Parse versions, skip pre-releases, drafts, and anything below MIN_VERSION
versions = []
for r in releases:
    if r["prerelease"] or r["draft"]:
        continue
    tag = r["tag_name"].lstrip("v")
    parts = tag.split(".")
    if len(parts) == 3 and all(p.isdigit() for p in parts):
        v = tuple(int(x) for x in parts)
        if v >= MIN_VERSION:
            versions.append(v)

# ── Group patches by (major, minor) ───────────────────────────────
by_minor = defaultdict(list)
for v in versions:
    by_minor[(v[0], v[1])].append(v)

result = []
if by_minor:
    latest_minor = max(by_minor)

    for minor in sorted(by_minor):
        patches = sorted(by_minor[minor])
        if minor == latest_minor:
            # Latest minor: every patch
            selected = patches
        else:
            # Older minors: only the newest patch
            selected = [patches[-1]]
        result.extend(".".join(map(str, v)) for v in selected)

min_str = ".".join(map(str, MIN_VERSION))
min_minor_prefix = f"{MIN_VERSION[0]}.{MIN_VERSION[1]}."
if not any(r.startswith(min_minor_prefix) for r in result):
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