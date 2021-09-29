-- MariaDB dump 10.19  Distrib 10.5.9-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: logs
-- ------------------------------------------------------
-- Server version	10.5.9-MariaDB-1:10.5.9+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sorted_logs`
--

DROP TABLE IF EXISTS `sorted_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sorted_logs` (
  `SLogID` int(11) NOT NULL AUTO_INCREMENT,
  `Test_LogID` int(11) NOT NULL,
  `SBlockId` varchar(128) NOT NULL,
  `Sender` varchar(128) NOT NULL,
  `Receiver` varchar(128) NOT NULL,
  `RequestDelay` int(11) DEFAULT NULL,
  `BlockDelay` int(11) DEFAULT NULL,
  `BlockDelivered` tinyint(1) DEFAULT NULL,
  `SDuplicate` tinyint(1) NOT NULL,
  PRIMARY KEY (`SLogID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sorted_logs`
--

LOCK TABLES `sorted_logs` WRITE;
/*!40000 ALTER TABLE `sorted_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sorted_logs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_logs`
--

DROP TABLE IF EXISTS `test_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `test_logs` (
  `LogID` int(11) NOT NULL AUTO_INCREMENT,
  `BlockId` varchar(128) NOT NULL,
  `LocalPeer` varchar(128) NOT NULL,
  `RemotePeer` varchar(128) NOT NULL,
  `SentAt` datetime(3) NULL DEFAULT NULL,
  `ReceivedAt` datetime(3) NULL DEFAULT NULL,
  `BlockRequestedAt` datetime(3) NULL DEFAULT NULL,
  `Duplicate` tinyint(1) NOT NULL,
  PRIMARY KEY (`LogID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_logs`
--

LOCK TABLES `sorted_logs` WRITE;
/*!40000 ALTER TABLE `sorted_logs` DISABLE KEYS */;
/*!40000 ALTER TABLE `sorted_logs` ENABLE KEYS */;
UNLOCK TABLES;


/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-16 15:50:24
