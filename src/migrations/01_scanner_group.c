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

#define _GNU_SOURCE

#include <stdio.h>

#include "declared.h"

/**
 * Add all active/admin users into the scanner group
 */
bool qol_migration_01_scanner_group(QolContext* context) {
	return qol_migration_push_active_admin_group(context, "scanner");
}
