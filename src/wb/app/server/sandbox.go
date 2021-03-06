package server

import (
	"regexp"
)

var blackList map[string]*regexp.Regexp = map[string]*regexp.Regexp{}

func addRule(key string) {
	blackList[key] = regexp.MustCompile("\\s" + key + "\\s")
}

func IsSandboxed(program string) (string, bool) {
	for key, re := range blackList {
		if re.FindString(program) != "" {
			return key, true
		}
	}
	return "", false
}

func init() {
	addRule("exit")
	addRule("setjmp")
	addRule("longjmp")
	addRule("signal")
	addRule("raise")
	addRule("asctime")
	addRule("ctime")
	addRule("difftime")
	addRule("gmtime")
	addRule("localtime")
	addRule("mktime")
	addRule("strftime")
	addRule("memchr")
	addRule("memmove")
	addRule("memset")
	addRule("strcat")
	addRule("strncat")
	addRule("strchr")
	addRule("strcmp")
	addRule("strncmp")
	addRule("strcoll")
	addRule("strcpy")
	addRule("strncpy")
	addRule("strcspn")
	addRule("staddRulerror")
	addRule("strlen")
	addRule("strpbrk")
	addRule("strrchr")
	addRule("strspn")
	addRule("strstr")
	addRule("strtok")
	addRule("strxfrm")
	addRule("abort")
	addRule("atexit")
	addRule("bsearch")
	addRule("calloc")
	addRule("getenv")
	addRule("mblen")
	addRule("mbstowcs")
	addRule("mbtowc")
	addRule("qsort")
	addRule("addRulealloc")
	addRule("srand")
	addRule("strtod")
	addRule("strtol")
	addRule("strtoul")
	addRule("system")
	addRule("wcstombs")
	addRule("wctomb")
	addRule("cleaaddRulerr")
	addRule("fclose")
	addRule("feof")
	addRule("ferror")
	addRule("fflush")
	addRule("fgetpos")
	addRule("fopen")
	addRule("faddRulead")
	addRule("faddRuleopen")
	addRule("fseek")
	addRule("fsetpos")
	addRule("ftell")
	addRule("fwrite")
	addRule("addRulemove")
	addRule("addRulename")
	addRule("addRulewind")
	addRule("setbuf")
	addRule("setvbuf")
	addRule("tmpfile")
	addRule("tmpnam")
	addRule("fprintf")
	addRule("fscanf")
	addRule("scanf")
	addRule("sprintf")
	addRule("sscanf")
	addRule("vfprintf")
	addRule("vprintf")
	addRule("vsprintf")
	addRule("fgetc")
	addRule("fgets")
	addRule("fputc")
	addRule("fputs")
	addRule("getc")
	addRule("getchar")
	addRule("gets")
	addRule("putc")
	addRule("putchar")
	addRule("puts")
	addRule("ungetc")
	addRule("perror")
	addRule("stdout")
	addRule("stderr")
	addRule("stdin")
	addRule("va_start")
	addRule("va_arg")
	addRule("va_end")
	addRule("__asm__")
}
