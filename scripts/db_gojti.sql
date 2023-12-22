-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 22, 2023 at 09:33 AM
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
  `number` varchar(16) NOT NULL,
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
(1, '081291808448', 2, 1, '2023-12-21 13:42:24', 'System', '2023-12-22 04:31:35', 'Muhamad Ulil Absor', 0),
(2, '081591808441', 1, 2, '2023-12-21 13:42:24', 'System', '2023-12-21 15:57:05', 'Muhamad Ulil Absor', 0),
(3, '+6281291808448', 2, 1, '2023-12-22 00:09:21', 'Muhamad Ulil Absor', NULL, NULL, 0),
(4, '6281291808448', 2, 1, '2023-12-22 00:10:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(5, '081291808447', 1, 1, '2023-12-22 00:19:39', 'Muhamad Ulil Absor', NULL, NULL, 0),
(6, '081203214895', 1, 1, '2023-12-22 03:14:55', 'Muhamad Ulil Absor', NULL, NULL, 0),
(7, '081203215849', 1, 1, '2023-12-22 03:30:49', 'Muhamad Ulil Absor', NULL, NULL, 0),
(8, '085703215868', 2, 2, '2023-12-22 03:31:08', 'Muhamad Ulil Absor', NULL, NULL, 0),
(9, '081203215877', 1, 1, '2023-12-22 03:31:17', 'Muhamad Ulil Absor', NULL, NULL, 0),
(10, '081703215885', 1, 3, '2023-12-22 03:31:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(11, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(12, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(13, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(14, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(15, '081703216920', 2, 3, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(16, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(17, '081703216920', 2, 3, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(18, '088103216920', 2, 5, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(19, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(20, '081203216920', 2, 1, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(21, '088103216920', 2, 5, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(22, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(23, '088103216920', 2, 5, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(24, '081203216920', 2, 1, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(25, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(26, '081703216920', 2, 3, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(27, '089603216920', 2, 4, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(28, '088103216920', 2, 5, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(29, '081203216920', 2, 1, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(30, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(31, '081703216920', 2, 3, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(32, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(33, '081203216920', 2, 1, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(34, '085703216920', 2, 2, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(35, '081703216920', 2, 3, '2023-12-22 03:48:40', 'Muhamad Ulil Absor', NULL, NULL, 0),
(36, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(37, '081203217431', 1, 1, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(38, '081203217431', 1, 1, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(39, '089603217431', 1, 4, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(40, '088103217431', 1, 5, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(41, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(42, '085703217431', 1, 2, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(43, '089603217431', 1, 4, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(44, '088103217431', 1, 5, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(45, '089603217431', 1, 4, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(46, '085703217431', 1, 2, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(47, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(48, '088103217431', 1, 5, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(49, '081203217431', 1, 1, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(50, '081203217431', 1, 1, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(51, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(52, '085703217431', 1, 2, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(53, '085703217431', 1, 2, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(54, '088103217431', 1, 5, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(55, '081203217431', 1, 1, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(56, '085703217431', 1, 2, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(57, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(58, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(59, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(60, '081703217431', 1, 3, '2023-12-22 03:57:11', 'Muhamad Ulil Absor', NULL, NULL, 0),
(61, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(62, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(63, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(64, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(65, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(66, '089603217445', 1, 4, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(67, '085703217445', 1, 2, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(68, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(69, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(70, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(71, '089603217445', 1, 4, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(72, '089603217445', 1, 4, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(73, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(74, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(75, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(76, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(77, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(78, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(79, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(80, '089603217445', 1, 4, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(81, '081703217445', 1, 3, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(82, '085703217445', 1, 2, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(83, '088103217445', 1, 5, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(84, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(85, '081203217445', 1, 1, '2023-12-22 03:57:25', 'Muhamad Ulil Absor', NULL, NULL, 0),
(86, '081291808445', 1, 1, '2023-12-22 04:19:48', 'Muhamad Ulil Absor', NULL, NULL, 0),
(87, '088103218798', 2, 5, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(88, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(89, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(90, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(91, '088103218798', 2, 5, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(92, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(93, '081703218798', 2, 3, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(94, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(95, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(96, '089603218798', 2, 4, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(97, '081703218798', 2, 3, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(98, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(99, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(100, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(101, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(102, '081703218798', 2, 3, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(103, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(104, '088103218798', 2, 5, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(105, '089603218798', 2, 4, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(106, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(107, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(108, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(109, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(110, '081203218798', 2, 1, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(111, '085703218798', 2, 2, '2023-12-22 04:19:58', 'Muhamad Ulil Absor', NULL, NULL, 0),
(112, '081291808425', 1, 1, '2023-12-22 08:26:45', 'Muhamad Ulil Absor', NULL, NULL, 0);

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

INSERT INTO `user` (`id`, `name`, `email`, `password`, `user_contact_id`, `role_id`, `active`, `created_at`, `created_by`, `modified_at`, `modified_by`, `is_deleted`) VALUES
(1, 'Muhamad Ulil Absor', 'absoralvord07@gmail.com', 'bWVyYWlobWltcGk=', 1, 1, 1, '2022-11-19 08:37:52', 'SYSTEM', NULL, NULL, 0);

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
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=113;

--
-- AUTO_INCREMENT for table `provider`
--
ALTER TABLE `provider`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

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
