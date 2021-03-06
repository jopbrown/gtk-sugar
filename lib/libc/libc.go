package libc

import sugar "github.com/jopbrown/gtk-sugar"

// FUNCTION_NAME = malloc, NONE, POINTER, 1, INT
func Malloc(size int) sugar.CandyWrapper {
	id := Candy().Guify("malloc", size).String()
	return Candy().NewWrapper(id)
}

// FUNCTION_NAME = calloc, NONE, POINTER, 2, INT, INT

// FUNCTION_NAME = free, NONE, NONE, 1, POINTER
func Free(ptr sugar.CandyWrapper) string {
	return Candy().Guify("free", ptr).String()
}

// FUNCTION_NAME = memset, NONE, NONE, 3, POINTER, INT, INT
func Memset(ptr sugar.CandyWrapper, val, size int) string {
	return Candy().Guify("memset", ptr, val, size).String()
}

// FUNCTION_NAME = printf, NONE, NONE, 2, STRING, POINTER
// FUNCTION_NAME = snprintf, NONE, NONE, 4, POINTER, INT, STRING, DOUBLE
// FUNCTION_NAME = strncpy, NONE, STRING, 3, POINTER, POINTER, INT
// FUNCTION_NAME = sscanf, NONE, INT, 3, POINTER, STRING, PTR_INT
// FUNCTION_NAME = strtok, NONE, STRING, 2, STRING, STRING
// FUNCTION_NAME = strsep, NONE, STRING, 2, LONG, STRING
// FUNCTION_NAME = strlen, NONE, INT, 1, POINTER

// FUNCTION_NAME = abs, NONE, INT, 1, INT
// FUNCTION_NAME = toascii, NONE, INT, 1, INT
// FUNCTION_NAME = putchar, NONE, INT, 1, STRING
// FUNCTION_NAME = kill, NONE, INT, 2, INT, INT
// FUNCTION_NAME = write, NONE, INT, 3, INT, STRING, INT
// FUNCTION_NAME = fsync, NONE, INT, 1, INT
// FUNCTION_NAME = rand, NONE, INT, 0
// FUNCTION_NAME = random, NONE, INT, 0
// FUNCTION_NAME = srand, NONE, NONE, 1, INT
// FUNCTION_NAME = srandom, NONE, NONE, 1, INT
// FUNCTION_NAME = time, NONE, INT, 1, NULL

// FUNCTION_NAME = div, NONE, INT, 2, INT, INT
// FUNCTION_NAME = fmod, NONE, DOUBLE, 2, DOUBLE, DOUBLE
// FUNCTION_NAME = fma, NONE, DOUBLE, 3, DOUBLE, DOUBLE, DOUBLE
// FUNCTION_NAME = usleep, NONE, NONE, 1, INT
