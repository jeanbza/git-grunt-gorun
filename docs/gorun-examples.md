# Usage Examples

## Basic running

```javascript
gorun: {
  first: {
    src: "main.go"
  }
}
```

## Running with flags

```javascript
gorun: {
  first: {
    src: "main.go"
  },
  flags: {
    boom: "bam",
    foo: 5
  }
}
```