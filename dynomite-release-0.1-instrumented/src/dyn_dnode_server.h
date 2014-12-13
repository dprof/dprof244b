/*
 * Dynomite - A thin, distributed replication layer for multi non-distributed storages.
 * Copyright (C) 2014 Netflix, Inc.
 */ 

#ifndef _DYN_DNODE_SERVER_H_
#define _DYN_DNODE_SERVER_H_

#include "dyn_core.h"

void dnode_ref(struct conn *conn, void *owner);
void dnode_unref(struct conn *conn);
void dnode_close(struct context *ctx, struct conn *conn);

rstatus_t dnode_each_init(void *elem, void *data);
rstatus_t dnode_each_deinit(void *elem, void *data);

rstatus_t dnode_init(struct context *ctx);
void dnode_deinit(struct context *ctx);
rstatus_t dnode_recv(struct context *ctx, struct conn *conn);

#endif

