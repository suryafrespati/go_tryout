module dsa

go 1.15

replace stack => ./stack

replace array_slice_example => ./array_slice_example

replace map_example => ./map_example

replace func_example => ./func_example

replace pointer_example => ./pointer_example

replace struct_example => ./struct_example

replace embedded_struct_example => ./embedded_struct_example

replace method_example => ./method_example

replace interface_example => ./interface_example

replace reflect_example => ./reflect_example

replace goroutine_example => ./goroutine_example

require (
	array_slice_example v0.0.0-00010101000000-000000000000
	embedded_struct_example v0.0.0-00010101000000-000000000000
	func_example v0.0.0-00010101000000-000000000000
	goroutine_example v0.0.0-00010101000000-000000000000
	interface_example v0.0.0-00010101000000-000000000000
	map_example v0.0.0-00010101000000-000000000000
	method_example v0.0.0-00010101000000-000000000000
	pointer_example v0.0.0-00010101000000-000000000000
	reflect_example v0.0.0-00010101000000-000000000000
	stack v0.0.0-00010101000000-000000000000
	struct_example v0.0.0-00010101000000-000000000000
)
