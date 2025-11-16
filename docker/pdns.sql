-- Copyright (c) HashiCorp, Inc.
-- SPDX-License-Identifier: MPL-2.0

/*
curl -X POST \
  'http://localhost:8081/api/v1/servers/localhost/zones' \
  --header 'X-API-Key: npm' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "name": "example.com.",
  "kind": "Native",
  "masters": [],
  "nameservers": [
    "ns1.pdns.",
    "ns2.pdns."
  ]
}'
*/

/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-12.0.2-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: pdns
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
-- Table structure for table `comments`
--

DROP TABLE IF EXISTS `comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `comments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `type` varchar(10) NOT NULL,
  `modified_at` int(11) NOT NULL,
  `account` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_uca1400_ai_ci DEFAULT NULL,
  `comment` text CHARACTER SET utf8mb3 COLLATE utf8mb3_uca1400_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `comments_name_type_idx` (`name`,`type`),
  KEY `comments_order_idx` (`domain_id`,`modified_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comments`
--

LOCK TABLES `comments` WRITE;
/*!40000 ALTER TABLE `comments` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `comments` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `cryptokeys`
--

DROP TABLE IF EXISTS `cryptokeys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `cryptokeys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) NOT NULL,
  `flags` int(11) NOT NULL,
  `active` tinyint(1) DEFAULT NULL,
  `published` tinyint(1) DEFAULT 1,
  `content` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `domainidindex` (`domain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cryptokeys`
--

LOCK TABLES `cryptokeys` WRITE;
/*!40000 ALTER TABLE `cryptokeys` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `cryptokeys` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `domainmetadata`
--

DROP TABLE IF EXISTS `domainmetadata`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `domainmetadata` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) NOT NULL,
  `kind` varchar(32) DEFAULT NULL,
  `content` text DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `domainmetadata_idx` (`domain_id`,`kind`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `domainmetadata`
--

LOCK TABLES `domainmetadata` WRITE;
/*!40000 ALTER TABLE `domainmetadata` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `domainmetadata` VALUES
(1,1,'SOA-EDIT-API','DEFAULT');
/*!40000 ALTER TABLE `domainmetadata` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `domains`
--

DROP TABLE IF EXISTS `domains`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `domains` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `master` varchar(128) DEFAULT NULL,
  `last_check` int(11) DEFAULT NULL,
  `type` varchar(8) NOT NULL,
  `notified_serial` int(10) unsigned DEFAULT NULL,
  `account` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_uca1400_ai_ci DEFAULT NULL,
  `options` varchar(64000) DEFAULT NULL,
  `catalog` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_index` (`name`),
  KEY `catalog_idx` (`catalog`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `domains`
--

LOCK TABLES `domains` WRITE;
/*!40000 ALTER TABLE `domains` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `domains` VALUES
(1,'example.com','',NULL,'NATIVE',NULL,'',NULL,NULL);
/*!40000 ALTER TABLE `domains` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `records`
--

DROP TABLE IF EXISTS `records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `domain_id` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  `content` varchar(64000) DEFAULT NULL,
  `ttl` int(11) DEFAULT NULL,
  `prio` int(11) DEFAULT NULL,
  `disabled` tinyint(1) DEFAULT 0,
  `ordername` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin DEFAULT NULL,
  `auth` tinyint(1) DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `nametype_index` (`name`,`type`),
  KEY `domain_id` (`domain_id`),
  KEY `ordername` (`ordername`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `records`
--

LOCK TABLES `records` WRITE;
/*!40000 ALTER TABLE `records` DISABLE KEYS */;
set autocommit=0;
INSERT INTO `records` VALUES
(1,1,'example.com','NS','ns1.pdns',1500,0,0,NULL,1),
(2,1,'example.com','NS','ns2.pdns',1500,0,0,NULL,1),
(12,1,'example.com','SOA','a.misconfigured.dns.server.invalid hostmaster.example.com 2025111205 10800 3600 604800 3600',1500,0,0,NULL,1);
/*!40000 ALTER TABLE `records` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `supermasters`
--

DROP TABLE IF EXISTS `supermasters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `supermasters` (
  `ip` varchar(64) NOT NULL,
  `nameserver` varchar(255) NOT NULL,
  `account` varchar(40) CHARACTER SET utf8mb3 COLLATE utf8mb3_uca1400_ai_ci NOT NULL,
  PRIMARY KEY (`ip`,`nameserver`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `supermasters`
--

LOCK TABLES `supermasters` WRITE;
/*!40000 ALTER TABLE `supermasters` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `supermasters` ENABLE KEYS */;
UNLOCK TABLES;
commit;

--
-- Table structure for table `tsigkeys`
--

DROP TABLE IF EXISTS `tsigkeys`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `tsigkeys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `algorithm` varchar(50) DEFAULT NULL,
  `secret` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `namealgoindex` (`name`,`algorithm`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tsigkeys`
--

LOCK TABLES `tsigkeys` WRITE;
/*!40000 ALTER TABLE `tsigkeys` DISABLE KEYS */;
set autocommit=0;
/*!40000 ALTER TABLE `tsigkeys` ENABLE KEYS */;
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

-- Dump completed on 2025-11-12 14:49:50
