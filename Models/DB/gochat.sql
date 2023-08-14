/*
 Navicat Premium Data Transfer

 Source Server         : gochat
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : gochat

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 03/06/2023 20:30:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for jwt_tokens
-- ----------------------------
DROP TABLE IF EXISTS `jwt_tokens`;
CREATE TABLE `jwt_tokens`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'jwt id',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户id',
  `header` text CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT '存储JWT的头部',
  `payload` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '存储JWT的有效载荷',
  `signature` varchar(255) CHARACTER SET ascii COLLATE ascii_bin NOT NULL COMMENT '存储JWT的签名',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '存储JWT的创建时间，默认为当前时间',
  `expires_at` timestamp NULL DEFAULT NULL COMMENT '存储JWT的过期时间，可以为空',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for message_table
-- ----------------------------
DROP TABLE IF EXISTS `message_table`;
CREATE TABLE `message_table`  (
  `record_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '单条对话记录id',
  `chat_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '对话id',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户id',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `role` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '对话角色',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '对话文本内容',
  `prompt_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT 'prompt令牌数',
  `completion_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT '回复令牌数',
  `total_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT '总令牌数',
  PRIMARY KEY (`record_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for message_table_bak
-- ----------------------------
DROP TABLE IF EXISTS `message_table_bak`;
CREATE TABLE `message_table_bak`  (
  `chat_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '对话id',
  `user_id` int UNSIGNED NOT NULL COMMENT '用户id',
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `role` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '对话角色',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '对话文本内容',
  `prompt_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT 'prompt令牌数',
  `completion_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT '回复令牌数',
  `total_tokens` int UNSIGNED NULL DEFAULT NULL COMMENT '总令牌数',
  PRIMARY KEY (`chat_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `pwd` varchar(15) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户密码',
  `admin` int(1) UNSIGNED ZEROFILL NOT NULL DEFAULT 0 COMMENT '管理员权限',
  `name` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户名',
  `api` varchar(51) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户存储的api',
  `api_status` int(1) UNSIGNED ZEROFILL NOT NULL DEFAULT 0 COMMENT '用户api是否可用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = COMPACT;

SET FOREIGN_KEY_CHECKS = 1;
