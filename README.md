# Noteworthy

Noteworthy is a personalized Bear note Golang app which configures Forever Notes,
Google Calendar etc.

## Goals

### Forever Notes

- In contents tags enclosed in `<>` should be replaced with their respective values
- Month should be in short word ex Jan, Feb

#### The Year Note

- A "YEAR" note that is sync'ed with all month and quarter notes. Ex: "2025"
- The contents of the "YEAR" note should be like below:

```
# <YEAR>

## Calendar
[[Q1 <YEAR>]] — [[Jan <YEAR>]] · [[Feb <YEAR>]] · [[Mar <YEAR>]]
[[Q2 <YEAR>]] — [[Apr <YEAR>]] · [[May <YEAR>]] · [[Jun <YEAR>]]
[[Q3 <YEAR>]] — [[Jul <YEAR>]] · [[Aug <YEAR>]] · [[Sep <YEAR>]]
[[Q4 <YEAR>]] — [[Oct <YEAR>]] · [[Nov <YEAR>]] · [[Dec <YEAR>]]

## Goals
- [ ]

## Projects
*

## Areas
*

- - -
#<YEAR>
```

- The Year note should be automatically created for current year along with child notes
- The Projects and Areas section should be left untouched for any changes in other sections

#### The Quarter Note

- Titled as "Q<#> YEAR". Ex: "Q1 2025"
- The contents will be:

```
# Q<#> <YEAR>

[[<YEAR>]] · ← [[Q<PREV> <YEAR>]] · [[Q<NEXT> <YEAR>]] →
[[<MONTH1> <YEAR>]] · [[<MONTH2> <YEAR>]] · [[<MONTH3> <YEAR>]]

## Theme
*

## Goals
- [ ]


- - -
#<YEAR>/Q<#>
```

#### The Monthly Note

- Titled as "<MONTH> YEAR". Ex: "Jan 2025"
- Contents below:

```
# <MONTH> <YEAR>

[[<YEAR>]] · [[Q<CURRENTQ> <YEAR>]] · ← [[<MONTHPREV> <YEAR>]] · [[<MONTHNEXT> <YEAR>]] →
Week <#> — [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]]
Week <#> — [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]]
Week <#> — [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]]
Week <#> — [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]]

## Goals
- [ ]

## Highlights
-

- - -
#<YEAR>/Q<#>/<MONTH>
```

#### The Daily Note

- Titled as "<DATE> <MONTH> <YEAR>". Ex: "23 Jan 2025"
- Contents below:

```
# <DATE> <MONTH> <YEAR>

[[<YEAR>]] · [[Q<CURRENTQ> <YEAR>]] · [[<MONTH> <YEAR>]] · ← [[<DATE> <MONTH> <YEAR>]] · [[<DATE> <MONTH> <YEAR>]] →

## Goals
- [ ]

## Highlights
-

## Events
### <HH:MM> → <HH:MM> — <MEETING TITLE>
- Attendees: <NAME1>, <NAME2>
- Location: <LOCATION>


- - -
#<YEAR>/Q<#>/<MONTH>/<DATE>
```
