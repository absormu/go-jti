-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 26, 2023 at 04:34 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_gojti`
--

-- --------------------------------------------------------

--
-- Table structure for table `phone_number`
--

CREATE TABLE `phone_number` (
  `id` bigint(20) NOT NULL,
  `number` varchar(100) NOT NULL,
  `type` tinyint(1) NOT NULL,
  `provider_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` varchar(100) NOT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  `modified_by` varchar(100) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `phone_number`
--

INSERT INTO `phone_number` (`id`, `number`, `type`, `provider_id`, `created_at`, `created_by`, `modified_at`, `modified_by`, `is_deleted`) VALUES
(1, 'fcY5IeFkf36lvEiSIiysLw==', 1, 3, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(2, '3j+ers0T9tP6GJNwpyFBHw==', 1, 1, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(3, 'k3qeH5vRJV6mpnkDg07cZg==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(4, 'QYmh2YP/zoqH5Ca3iRvO0g==', 1, 1, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(5, 'WzhYYNveJ+eaqiSlo6TScA==', 1, 2, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(6, 'Am+4Tf6LwiDEUdDysWxsEA==', 2, 5, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(7, '+5sNXrZVM06KRicnVtHAdg==', 1, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(8, 'sYyGDN964Rack886Wfereg==', 2, 5, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(9, 'Iow6mQOSUdqaQFssomk1sw==', 1, 5, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(10, 'R9ZkO7K/kVis+5tVCBK79w==', 2, 2, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(11, '7MyAzgxYjzczCaM7kSgEfw==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(12, 'xf6uBggWERl6nquNjsvpfQ==', 1, 3, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(13, 'cgoPf7wglhGB0FoHoe1QEQ==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(14, 'H39f1asZpgk6BhYC/t8+zw==', 1, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(15, 'e2pql/bsVJ+CT/nSxeLVSQ==', 1, 1, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(16, 'kDeXWq66O2LzzviyNCFAdg==', 1, 1, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(17, 'NDaBsKBpg/rpG1Q86nsYdA==', 1, 1, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(18, 'wZH41cVQHOwX08EtqH57bQ==', 1, 2, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(19, 'a3ooDVVrNrjUPEMF/T2I2A==', 2, 2, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(20, '03kdoST0ogMPEwHE4DdueA==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(21, 'Pvzoai6HYZNfbPHSjMiEAA==', 1, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(22, 'xoy+ey7Ge5ONlF1e3Yrisg==', 2, 2, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(23, '4RyZq83w5SuMIvRY2yUQfA==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(24, 'e7pWfzeSEjtWWLRjUzN1Vg==', 2, 5, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0),
(25, 'pXvAwOafDbpdx7cCkQ6RWg==', 2, 4, '2023-12-26 15:30:20', 'absormu', NULL, NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `provider`
--

CREATE TABLE `provider` (
  `id` bigint(20) NOT NULL,
  `code` varchar(5) NOT NULL,
  `name` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `provider`
--

INSERT INTO `provider` (`id`, `code`, `name`) VALUES
(1, '0812', 'TELKOMSEL'),
(2, '0857', 'INDOSAT'),
(3, '0817', 'XL'),
(4, '0896', 'TRI'),
(5, '0881', 'SMARTFEN');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(150) NOT NULL,
  `password` varchar(150) NOT NULL,
  `user_contact_id` bigint(20) NOT NULL,
  `role_id` int(11) NOT NULL,
  `token` text DEFAULT NULL,
  `active` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `created_by` varchar(100) NOT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  `modified_by` varchar(100) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `user_contact_id`, `role_id`, `token`, `active`, `created_at`, `created_by`, `modified_at`, `modified_by`, `is_deleted`) VALUES
(1, 'Muhamad Ulil Absor', 'absoralvord07@gmail.com', 'bWVyYWlobWltcGk=', 1, 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNzAzNjkwODQzLCJuYW1lIjoiYWJzb3JtdSIsInVpZCI6ImNtNWYxbXUzbnQ0a3MwMWJsaGxnIiwidXNlcl9pZCI6IjEwNDkzMjg2MTU3MDkxNDM5NzgyMSJ9.GNgTzc9CI4oCw4gKrmK_dmDLOJZfmF4f3tlmKZJTpyM', 1, '2022-11-19 08:37:52', 'SYSTEM', NULL, NULL, 0),
(2, 'absormu', 'absormu@gmail.com', '104932861570914397821', 0, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yYWx2b3JkMDdAZ21haWwuY29tIiwiZXhwIjoxNzAzNjg3Mzg1LCJuYW1lIjoiYWJzb3JtdSIsInVpZCI6ImNtNWU2bWUzbnQ0a2M5NHY5YTUwIiwidXNlcl9pZCI6IjEwNDkzMjg2MTU3MDkxNDM5NzgyMSJ9.ks0g4zVREHQgk4dtrys7oeD-CjhlVs4A1LSnp6DnADI', 1, '2023-12-26 14:29:45', '', NULL, NULL, 0),
(3, 'absor developers', 'absordev9@gmail.com', '102069942941764343311', 0, 0, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFic29yZGV2OUBnbWFpbC5jb20iLCJleHAiOjE3MDM2OTExMTQsIm5hbWUiOiJhYnNvciBkZXZlbG9wZXJzIiwidWlkIjoiY201ZjNxbTNudDRrczAxYmxpa2ciLCJ1c2VyX2lkIjoiMTAyMDY5OTQyOTQxNzY0MzQzMzExIn0.gaIqAemDkcn29k95betFBN5i4WALds0SKNdm20dlnrs', 1, '2023-12-26 15:30:50', '', NULL, NULL, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `phone_number`
--
ALTER TABLE `phone_number`
  ADD PRIMARY KEY (`id`),
  ADD KEY `provider_id` (`provider_id`);

--
-- Indexes for table `provider`
--
ALTER TABLE `provider`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `phone_number`
--
ALTER TABLE `phone_number`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `provider`
--
ALTER TABLE `provider`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `phone_number`
--
ALTER TABLE `phone_number`
  ADD CONSTRAINT `phone_number_ibfk_1` FOREIGN KEY (`provider_id`) REFERENCES `provider` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
