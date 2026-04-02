A map in Go is used to store data in key:value pairs.
Think of it like a dictionary:

key → name
value → actual data

Example:

"name" → "John"
"age" → 25
* Key Points
A map stores data using keys and values
Each key must be unique (no duplicates)
The order is not fixed (it can change anytime)
You can add, update, delete, and read data easily
* Creating a Map

You can create a map:

Directly with values
Or using make() (best way for empty maps)
* Using a Map
Add or update:
map[key] = value
Get value:
map[key]
Delete:
delete(map, key)
Check if key exists:
value, ok := map[key]
* Important Things
If you don’t use make() for an empty map, it can crash
Maps don’t keep order
Some types (like slices, maps, functions) cannot be keys
* Looping Through a Map

You can use range to go through all items:

It gives you key and value
* Special Note

Maps are references, meaning:

If you copy a map and change one, the other also changes
