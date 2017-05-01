Sit down, please
================

Seating arrangements is hard to get right, so this takes a JSON
structure of the guests and arranges them automatically and naively
based on a few rules:

0. couples don't get seated next to each other
0. it's boy-girl-boy-girl
0. couples are on the same table, at least

Data
====

The structure is based around an array (of all units) of arrays (of each
unit). If the unit array is two, that means it's a couple. I do not
support bigamy atm.

```json
[
    [
        { "name": "Jane Doe", "gender": "female" },
        { "name": "John Doe", "gender": "male" }
    ]
]
```
