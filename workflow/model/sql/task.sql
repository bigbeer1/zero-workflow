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

 Date: 07/04/2022 13:54:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for task
-- 任务表
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务ID',
  `created_at` bigint(36) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `claim_time` bigint(36) NULL DEFAULT NULL COMMENT '最近通过时间',
  `node_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '当前执行流所在的节点ID',
  `node_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '步骤名',
  `node_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '步骤类型',
  `step` int(11) NOT NULL COMMENT '第几步',
  `procinst_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '流程实例id',
  `assignee_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '受理人ID',
  `assignee_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '受理人姓名',
  `un_complete_num` tinyint(3) NOT NULL DEFAULT 0 COMMENT '需要审批的数',
  `agree_num` tinyint(3) NOT NULL DEFAULT 0 COMMENT '审批通过数',
  `is_finished` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否完成',
  `tenant_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '租户ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `task_proc_inst_id_proc_inst_id_foreign`(`procinst_id`) USING BTREE,
  INDEX `idx_id`(`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
