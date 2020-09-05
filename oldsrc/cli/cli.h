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

/**
 * Basic CLI utilities
 */
typedef struct SubCommand {
	const char* name; /**<Name of this subcommand */
	bool (*execute)(int argc, char** argv);

	const char* short_desc; /**<Short description to display */
} SubCommand;

/**
 * List users on the system
 */
bool qol_cli_list_users(int argc, char** argv);

/**
 * Perform any migration tasks as deemed needed
 */
bool qol_cli_migrate(int argc, char** argv);

/**
 * Trigger a migration on next boot
 */
bool qol_cli_trigger(int argc, char** argv);
