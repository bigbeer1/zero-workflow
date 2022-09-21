/*
 Navicat Premium Data Transfer

 Source Server         : dabenxiong
 Source Server Type    : MySQL
 Source Server Version : 100424
 Source Host           : localhost:3306
 Source Schema         : workflow

 Target Server Type    : MySQL
 Target Server Version : 100424
 File Encoding         : 65001

 Date: 08/04/2022 17:09:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for execution
-- 流程实例执行表
-- ----------------------------
DROP TABLE IF EXISTS `execution`;
CREATE TABLE `execution`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '执行ID',
  `procinst_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '实例ID',
  `procdef_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '流程名',
  `node_infos` varbinary(4000) NOT NULL COMMENT '节点信息',
  `start_time` bigint(36) NOT NULL COMMENT '开始时间',
  `tenant_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `execution_proc_inst_id_proc_inst_id_foreign`(`procinst_id`) USING BTREE,
  INDEX `idx_id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
