# Golang memo

## Map vs Struct

### Map 
- All keys must be the same type
- All values must be the same type
- Keys are indexed - can iterate over them
- Use to represent a collection of related properties
- Don't need to know all the keys at complile time
- `Reference Type`

### Struct 
- Values can be of different type
- Keys don't support indexing
- `Value Type`
- Need to know all the different fields at compile time
- Use to represent a "thing" with a lot of different properties
