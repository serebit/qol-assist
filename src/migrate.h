/*
 * This file is part of qol-assist.
 *
 * Copyright © 2017-2018 Solus Project
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 */

#pragma once

#include <stdbool.h>

#include "user-manager.h"

/**
 * QolContext is a bundling type passed to all of our functions, which will
 * afford us the ability to grow in the future..
 */
typedef struct QolContext {
	QolUserManager* user_manager;
} QolContext;

/**
 * Create a new QolContext, which will initialise itself with all the
 * appropriate helpers.
 *
 * @returns A newly allocated QolContext
 */
QolContext* qol_context_new(void);

/**
 * Free an existing QolContext
 *
 * @param context Pointer to an allocated context
 */
void qol_context_free(QolContext* context);

/**
 * Migration functions follow a very strict API and are provided access to
 * a QolContext to perform an execution.
 *
 * If a migration func returns true, we'll bump the migration level, otherwise
 * it is considered a migration error.
 */
typedef bool (*qol_migration_func)(QolContext* context);

/**
 * Helps to describe a QolMigration for logging purposes
 */
typedef struct QolMigration {
	const char* name;
	qol_migration_func func;
} QolMigration;
