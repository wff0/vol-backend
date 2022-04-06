CREATE TABLE IF NOT EXISTS `event`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `title`       varchar(100)  NOT NULL COMMENT '活动标题',
    `location`    varchar(100)  NOT NULL COMMENT '活动地点',
    `user_id`     int(11) NOT NULL COMMENT '发起人',
    `description` varchar(1000) NOT NULL COMMENT '活动介绍',
    `status`      varchar(20)   NOT NULL COMMENT '活动进行状态 报名中/进行中/已结束',
    `max_num`     int(11) NOT NULL COMMENT '上限人数',
    `start_time`  varchar(100)    NULL COMMENT '活动开始时间',
    `end_time`    varchar(100)    NULL COMMENT '活动结束时间',
    `create_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='志愿活动表';

CREATE TABLE IF NOT EXISTS `event_apply`
(
    `id`           int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `event_id`     varchar(100)  NOT NULL COMMENT '活动的id',
    `user_id`      int(11) NOT NULL COMMENT '申请人的id',
    `description`  varchar(1000) NOT NULL COMMENT '活动介绍',
    `apply_status` varchar(20)   NOT NULL COMMENT '申请状态 /报名中/报名通过/报名失败',
    `create_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='志愿活动申请表';

CREATE TABLE IF NOT EXISTS `news`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `title`      varchar(100)  NOT NULL COMMENT '新闻标题',
    `body`       varchar(2000) NOT NULL COMMENT '新闻主体',
    `create_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='新闻表';

CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `username`   varchar(100) NOT NULL COMMENT '用户昵称',
    `password`   varchar(100) NOT NULL COMMENT '用户密码',
    `school`     varchar(100) COMMENT '用户学校',
    `class`      varchar(100) COMMENT '用户班级',
    `Role`       int(11) NOT NULL COMMENT '用户角色 //0 普通用户 1 管理员',
    `create_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  timestamp     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='志愿活动表';