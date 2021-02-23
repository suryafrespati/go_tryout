module dsa

go 1.15

replace stack => ./stack

replace array_slice_example => ./array_slice_example

replace map_example => ./map_example

replace func_example => ./func_example

require (
	array_slice_example v0.0.0-00010101000000-000000000000
	func_example v0.0.0-00010101000000-000000000000
	map_example v0.0.0-00010101000000-000000000000
	stack v0.0.0-00010101000000-000000000000
)
