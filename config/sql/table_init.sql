CREATE TABLE IF NOT EXISTS user (
    id  BIGINT UNSIGNED COMMENT 'PK',
    user_name varchar(255) UNIQUE NOT NULL DEFAULT '' COMMENT '用户名称，唯一',
    email varchar(255) UNIQUE NOT NULL DEFAULT '' COMMENT '用户邮箱，唯一',
    password_digest varchar(255) NOT NULL DEFAULT '' COMMENT '加密后的密码',
    gender TINYINT NOT NULL DEFAULT -1 COMMENT '用户性别',
    avatar varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像的相对路径',
    fans INT NOT NULL DEFAULT 0 COMMENT '用户粉丝数量',
    follows INT NOT NULL DEFAULT 0 COMMENT '用户关注数量',
    totp_enable BOOLEAN NOT NULL DEFAULT false COMMENT '用户是否开启totp',
    totp_secret varchar(255) NOT NULL DEFAULT '' COMMENT '用户totp密钥',
    create_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'item创建时间',
    update_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'item更新时间',
    delete_at BIGINT UNSIGNED DEFAULT NULL COMMENT 'item删除时间',
    PRIMARY KEY pk_user(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户基础表' ;


CREATE TABLE IF NOT EXISTS video (
    id BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'PK',
    uid BIGINT NOT NULL DEFAULT 0 COMMENT '上传视频者',
    video_url varchar(255) NOT NULL DEFAULT '' COMMENT  '视频url',
    cover_url varchar(255) NOT NULL DEFAULT '' COMMENT '封面url',
    intro varchar(255) NOT NULL DEFAULT '' COMMENT '简介',
    title varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    video_ext varchar(6) NOT NULL DEFAULT '' COMMENT '文件后缀，用于oss获取path',
    cover_ext varchar(6) NOT NULL DEFAULT '' COMMENT '文件后缀，用于oss获取path',
    stars INT NOT NULL DEFAULT 0 COMMENT '收藏数',
    likes INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    views INT NOT NULL DEFAULT 0 COMMENT '浏览数',
    create_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'item创建时间',
    update_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'item更新时间',
    delete_at BIGINT UNSIGNED DEFAULT NULL COMMENT 'item删除时间',
    PRIMARY KEY pk_video(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='视频信息表' ;


CREATE TABLE IF NOT EXISTS comment (
    id BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'PK',
    uid BIGINT NOT NULL DEFAULT 0 COMMENT '评论人id',
    vid BIGINT NOT NULL DEFAULT 0 COMMENT '视频id',
    root_id BIGINT NOT NULL DEFAULT 0 COMMENT '为0的话则是对视频的评论',
    reply_id BIGINT NOT NULL DEFAULT 0 COMMENT '回复id，为0则说明是对视频的评论',
    content varchar(255) NOT NULL DEFAULT '' COMMENT '评论内容',
    likes INT NOT NULL DEFAULT 0 COMMENT '点赞数',
    create_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'item创建时间',
    delete_at BIGINT UNSIGNED DEFAULT NULL COMMENT 'item删除时间',
    PRIMARY KEY pk_comment(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表' ;

CREATE TABLE IF NOT EXISTS user_video_likes (
    user_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'user_id',
    video_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'video_id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表' ;

CREATE TABLE IF NOT EXISTS user_comment_likes (
    user_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'user_id',
    comment_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'comment_id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表' ;

CREATE TABLE IF NOT EXISTS chat_messages (
    id BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'PK',
    uid BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'sender_id',
    receiver_id BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'receiver_id',
    content varchar(255) NOT NULL DEFAULT '' COMMENT 'message内容',
    create_at BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'message创建时间',
    delete_at BIGINT UNSIGNED DEFAULT NULL COMMENT 'message删除时间',
    have_read BOOLEAN NOT NULL DEFAULT false COMMENT '已读',
    PRIMARY KEY chat_message_pk(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表' ;