# lineCounter
It's a tool to count effective code lines.
Usage: lineCounter -ext <file ext> -verbose <level of verbose> <file list>
	file ext: the surfix of files that will be counted. Default is "go".
	level of verbose: the level of details that will be shown in the result. Default is 0.
Remarks:
	The tool works with C family syntax codes, includes C/C++, java, Go, etc.
	Non-effective lines, including comments, empty line,etc.,  are not counted.
Examples:
	lineCounter .
	lineCounter *
	lineCounter */*
	lineCounter -ext c *
	lineCounter -ext pp -verbose 2 .
