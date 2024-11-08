# SBVJ01-Reader
Reads and writes [**Starbound**](https://starbounder.org/Starbound) save files


# Example
Check out [all the examples](https://github.com/hollowness-inside/SBVJ01-Reader/tree/main/examples)


```go
// Read from file
sbvj, _ := sbvj.ReadFile("data/file.player")

// File Options
opts := sbvj.Options
_filename := opts.Name
_versioned := opts.Versioned
_version := opts.Version

// Content
content := sbvj.Content.Value.(types.SBVJMap)
movController := content["movementController"].Value.(types.SBVJMap)
facDir := movController["facingDirection"].Value.(string)

fmt.Println("Movement Controller:", movController)
fmt.Println("Player facing direction:", facDir)
```

Output is the following
```json
Movement Controller: {"position": [0.000000, 0.000000], "movingDirection": "right", "rotation": 0.000000, "crouching": false, "facingDirection": "right", "velocity": [0.000000, 0.000000]}
Player facing direction: right
```

# Converting to JSON
```go
// Read the file
sbvj, _ := sbvj.ReadFile("data/file.player")

// Extract the content
content := sbvj.Content

// Marshal to JSON
jsoned, _ := json.Marshal(content)

// Save the JSON
output, _ := os.Create("data/output.json")
output.Write(jsoned)
```