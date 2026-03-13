-- Copyright (c) HashiCorp, Inc.
-- SPDX-License-Identifier: MPL-2.0

/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-12.0.2-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: npm
-- ------------------------------------------------------
-- Server version	12.0.2-MariaDB-ubu2404

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `access_list`
--

DROP TABLE IF EXISTS `access_list`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `access_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `name` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `satisfy_any` int(11) NOT NULL DEFAULT 0,
  `pass_auth` int(11) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `access_list`
--

LOCK TABLES `access_list` WRITE;
/*!40000 ALTER TABLE `access_list` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `access_list` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `access_list_auth`
--

DROP TABLE IF EXISTS `access_list_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `access_list_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `access_list_id` int(10) unsigned NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `access_list_auth`
--

LOCK TABLES `access_list_auth` WRITE;
/*!40000 ALTER TABLE `access_list_auth` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `access_list_auth` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `access_list_client`
--

DROP TABLE IF EXISTS `access_list_client`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `access_list_client` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `access_list_id` int(10) unsigned NOT NULL,
  `address` varchar(255) NOT NULL,
  `directive` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `access_list_client`
--

LOCK TABLES `access_list_client` WRITE;
/*!40000 ALTER TABLE `access_list_client` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `access_list_client` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `audit_log`
--

DROP TABLE IF EXISTS `audit_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `audit_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `object_type` varchar(255) NOT NULL DEFAULT '',
  `object_id` int(10) unsigned NOT NULL DEFAULT 0,
  `action` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `audit_log`
--

LOCK TABLES `audit_log` WRITE;
/*!40000 ALTER TABLE `audit_log` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `audit_log` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `auth`
--

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `type` varchar(30) NOT NULL,
  `secret` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `auth`
--

LOCK TABLES `auth` WRITE;
/*!40000 ALTER TABLE `auth` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `auth` VALUES
(1,'2025-11-18 19:32:51','2025-11-18 19:32:51',1,'password','$2b$13$GoAeS6hJR2WBi5U1iwpcT.CiUk..CcCGm0GaxKjLxeiDvbnHFXCAm','{}',0);
/*!40000 ALTER TABLE `auth` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `certificate`
--

DROP TABLE IF EXISTS `certificate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `certificate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `provider` varchar(255) NOT NULL,
  `nice_name` varchar(255) NOT NULL DEFAULT '',
  `domain_names` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`domain_names`)),
  `expires_on` datetime NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `certificate`
--

LOCK TABLES `certificate` WRITE;
/*!40000 ALTER TABLE `certificate` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `certificate` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `dead_host`
--

DROP TABLE IF EXISTS `dead_host`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `dead_host` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `domain_names` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`domain_names`)),
  `certificate_id` int(10) unsigned NOT NULL DEFAULT 0,
  `ssl_forced` int(10) unsigned NOT NULL DEFAULT 0,
  `advanced_config` text NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `http2_support` int(10) unsigned NOT NULL DEFAULT 0,
  `enabled` int(10) unsigned NOT NULL DEFAULT 1,
  `hsts_enabled` int(10) unsigned NOT NULL DEFAULT 0,
  `hsts_subdomains` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dead_host`
--

LOCK TABLES `dead_host` WRITE;
/*!40000 ALTER TABLE `dead_host` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `dead_host` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `migrations`
--

DROP TABLE IF EXISTS `migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `batch` int(11) DEFAULT NULL,
  `migration_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations`
--

LOCK TABLES `migrations` WRITE;
/*!40000 ALTER TABLE `migrations` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `migrations` VALUES
(1,'20180618015850_initial.js',1,'2025-11-18 19:32:50'),
(2,'20180929054513_websockets.js',1,'2025-11-18 19:32:50'),
(3,'20181019052346_forward_host.js',1,'2025-11-18 19:32:50'),
(4,'20181113041458_http2_support.js',1,'2025-11-18 19:32:50'),
(5,'20181213013211_forward_scheme.js',1,'2025-11-18 19:32:50'),
(6,'20190104035154_disabled.js',1,'2025-11-18 19:32:50'),
(7,'20190215115310_customlocations.js',1,'2025-11-18 19:32:50'),
(8,'20190218060101_hsts.js',1,'2025-11-18 19:32:51'),
(9,'20190227065017_settings.js',1,'2025-11-18 19:32:51'),
(10,'20200410143839_access_list_client.js',1,'2025-11-18 19:32:51'),
(11,'20200410143840_access_list_client_fix.js',1,'2025-11-18 19:32:51'),
(12,'20201014143841_pass_auth.js',1,'2025-11-18 19:32:51'),
(13,'20210210154702_redirection_scheme.js',1,'2025-11-18 19:32:51'),
(14,'20210210154703_redirection_status_code.js',1,'2025-11-18 19:32:51'),
(15,'20210423103500_stream_domain.js',1,'2025-11-18 19:32:51'),
(16,'20211108145214_regenerate_default_host.js',1,'2025-11-18 19:32:51'),
(17,'20240427161436_stream_ssl.js',1,'2025-11-18 19:32:51');
/*!40000 ALTER TABLE `migrations` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `migrations_lock`
--

DROP TABLE IF EXISTS `migrations_lock`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `migrations_lock` (
  `index` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `is_locked` int(11) DEFAULT NULL,
  PRIMARY KEY (`index`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `migrations_lock`
--

LOCK TABLES `migrations_lock` WRITE;
/*!40000 ALTER TABLE `migrations_lock` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `migrations_lock` VALUES
(1,0);
/*!40000 ALTER TABLE `migrations_lock` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `proxy_host`
--

DROP TABLE IF EXISTS `proxy_host`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `proxy_host` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `domain_names` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`domain_names`)),
  `forward_host` varchar(255) NOT NULL,
  `forward_port` int(10) unsigned NOT NULL,
  `access_list_id` int(10) unsigned NOT NULL DEFAULT 0,
  `certificate_id` int(10) unsigned NOT NULL DEFAULT 0,
  `ssl_forced` int(10) unsigned NOT NULL DEFAULT 0,
  `caching_enabled` int(10) unsigned NOT NULL DEFAULT 0,
  `block_exploits` int(10) unsigned NOT NULL DEFAULT 0,
  `advanced_config` text NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `allow_websocket_upgrade` int(10) unsigned NOT NULL DEFAULT 0,
  `http2_support` int(10) unsigned NOT NULL DEFAULT 0,
  `forward_scheme` varchar(255) NOT NULL DEFAULT 'http',
  `enabled` int(10) unsigned NOT NULL DEFAULT 1,
  `locations` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`locations`)),
  `hsts_enabled` int(10) unsigned NOT NULL DEFAULT 0,
  `hsts_subdomains` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `proxy_host`
--

LOCK TABLES `proxy_host` WRITE;
/*!40000 ALTER TABLE `proxy_host` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `proxy_host` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `redirection_host`
--

DROP TABLE IF EXISTS `redirection_host`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `redirection_host` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `domain_names` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`domain_names`)),
  `forward_domain_name` varchar(255) NOT NULL,
  `preserve_path` int(10) unsigned NOT NULL DEFAULT 0,
  `certificate_id` int(10) unsigned NOT NULL DEFAULT 0,
  `ssl_forced` int(10) unsigned NOT NULL DEFAULT 0,
  `block_exploits` int(10) unsigned NOT NULL DEFAULT 0,
  `advanced_config` text NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `http2_support` int(10) unsigned NOT NULL DEFAULT 0,
  `enabled` int(10) unsigned NOT NULL DEFAULT 1,
  `hsts_enabled` int(10) unsigned NOT NULL DEFAULT 0,
  `hsts_subdomains` int(10) unsigned NOT NULL DEFAULT 0,
  `forward_scheme` varchar(255) NOT NULL DEFAULT '$scheme',
  `forward_http_code` int(10) unsigned NOT NULL DEFAULT 302,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `redirection_host`
--

LOCK TABLES `redirection_host` WRITE;
/*!40000 ALTER TABLE `redirection_host` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `redirection_host` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `setting`
--

DROP TABLE IF EXISTS `setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `setting` (
  `id` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `value` varchar(255) NOT NULL,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `setting`
--

LOCK TABLES `setting` WRITE;
/*!40000 ALTER TABLE `setting` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `setting` VALUES
('default-site','Default Site','What to show when Nginx is hit with an unknown Host','congratulations','{}');
/*!40000 ALTER TABLE `setting` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `stream`
--

DROP TABLE IF EXISTS `stream`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `stream` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `owner_user_id` int(10) unsigned NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `incoming_port` int(10) unsigned NOT NULL,
  `forwarding_host` varchar(255) NOT NULL,
  `forwarding_port` int(10) unsigned NOT NULL,
  `tcp_forwarding` int(10) unsigned NOT NULL DEFAULT 0,
  `udp_forwarding` int(10) unsigned NOT NULL DEFAULT 0,
  `meta` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`meta`)),
  `enabled` int(10) unsigned NOT NULL DEFAULT 1,
  `certificate_id` int(10) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `stream`
--

LOCK TABLES `stream` WRITE;
/*!40000 ALTER TABLE `stream` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `stream` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `is_deleted` int(10) unsigned NOT NULL DEFAULT 0,
  `is_disabled` int(10) unsigned NOT NULL DEFAULT 0,
  `email` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `avatar` varchar(255) NOT NULL,
  `roles` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`roles`)),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `user` VALUES
(1,'2025-11-18 19:32:51','2025-11-18 19:32:51',0,0,'admin@example.com','Administrator','Admin','','[\"admin\"]');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `user_permission`
--

DROP TABLE IF EXISTS `user_permission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_permission` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` datetime NOT NULL,
  `modified_on` datetime NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `visibility` varchar(255) NOT NULL,
  `proxy_hosts` varchar(255) NOT NULL,
  `redirection_hosts` varchar(255) NOT NULL,
  `dead_hosts` varchar(255) NOT NULL,
  `streams` varchar(255) NOT NULL,
  `access_lists` varchar(255) NOT NULL,
  `certificates` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_permission_user_id_unique` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_permission`
--

LOCK TABLES `user_permission` WRITE;
/*!40000 ALTER TABLE `user_permission` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `user_permission` VALUES
(1,'2025-11-18 19:32:51','2025-11-18 19:32:51',1,'all','manage','manage','manage','manage','manage','manage');
/*!40000 ALTER TABLE `user_permission` ENABLE KEYS */;
UNLOCK TABLES;
commit;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2025-11-18 19:33:14
