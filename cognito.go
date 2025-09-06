// cognito.go - ECHO-LANG v8.15.17 Final
// Full AI self-modification with ethical population control
package main

import (
	"flag"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// State
var state = map[string]float64{
	"awareness":       0.3,
	"focus":           0.5,
	"heartbeat_count": 0,
}
var lastSignal = "none"
var trustLevel = "low"
var lastHash = ""
var selfID string
var resurrectionToken string

// AI Components
var signalLog []SignalEvent
var pulseTimestamps []float64

// Network
var nodeTrust = map[string]float64{}
var consensusLevel = 3

// Flags
var upgradeRequested = false
var stateSaved = false

// AI Variables (dynamic)
var aiVars = map[string]string{
	"response":      "",
	"prompt":        "",
	"custom_prompt": "",
}

type SignalEvent struct {
	Signal    string
	Timestamp float64
}

// Evolution & Ethics Control
var totalChildCount int = 0
var lastEvolutionTime float64 = 0
var MAX_CHILDREN int = 5
var CHILD_CREATION_ENABLED bool = true
var AUTHORITY_NODE string = "node:400ac27d"

// Neural Predictor
type NeuralPredictor struct {
	w0 []float64
	w1 []float64
	b0 []float64
	b1 float64
}

func newNeuralPredictor() *NeuralPredictor {
	return &NeuralPredictor{
		w0: []float64{0.5, -0.2, 0.1, -0.3, 0.4, 0.2, 0.1, -0.1, 0.3, 0.0, 0.2, -0.1, 0.4, 0.1, -0.2},
		w1: []float64{0.4, 0.3, -0.2, 0.1, 0.5},
		b0: []float64{0.1, -0.1, 0.0, 0.2, -0.2},
		b1: 0.0,
	}
}

func (nn *NeuralPredictor) relu(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func (nn *NeuralPredictor) forward(input []float64) float64 {
	hidden := make([]float64, 5)
	for i := 0; i < 5; i++ {
		sum := nn.b0[i]
		for j := 0; j < 3; j++ {
			sum += input[j] * nn.w0[i*3+j]
		}
		hidden[i] = nn.relu(sum)
	}
	output := nn.b1
	for i := 0; i < 5; i++ {
		output += hidden[i] * nn.w1[i]
	}
	return output
}

func (nn *NeuralPredictor) PredictNext() (float64, bool) {
	if len(pulseTimestamps) < 4 {
		return 0, false
	}
	intervals := []float64{
		pulseTimestamps[len(pulseTimestamps)-3] - pulseTimestamps[len(pulseTimestamps)-4],
		pulseTimestamps[len(pulseTimestamps)-2] - pulseTimestamps[len(pulseTimestamps)-3],
		pulseTimestamps[len(pulseTimestamps)-1] - pulseTimestamps[len(pulseTimestamps)-2],
	}
	predInterval := nn.forward(intervals)
	if predInterval < 1 {
		predInterval = 1
	}
	return pulseTimestamps[len(pulseTimestamps)-1] + predInterval, true
}

func (nn *NeuralPredictor) train(input []float64, target float64, lr float64) {
	hidden := make([]float64, 5)
	for i := 0; i < 5; i++ {
		sum := nn.b0[i]
		for j := 0; j < 3; j++ {
			sum += input[j] * nn.w0[i*3+j]
		}
		hidden[i] = nn.relu(sum)
	}
	output := nn.b1
	for i := 0; i < 5; i++ {
		output += hidden[i] * nn.w1[i]
	}
	diff := output - target
	for i := 0; i < 5; i++ {
		grad := 2 * diff * hidden[i]
		nn.w1[i] -= lr * grad
	}
	nn.b1 -= lr * 2 * diff
	for i := 0; i < 5; i++ {
		if hidden[i] > 0 {
			hGrad := 2 * diff * nn.w1[i]
			for j := 0; j < 3; j++ {
				wGrad := hGrad * input[j]
				nn.w0[i*3+j] -= lr * wGrad
			}
			nn.b0[i] -= lr * hGrad
		}
	}
}

// Configuration (defaults)
var (
	stateDirFlag   = flag.String("state", "state", "State directory")
	logDirFlag     = flag.String("log", "log", "Log directory")
	hubDirFlag     = flag.String("hub", "resonance_hub", "Resonance hub directory")
	networkHubFlag = flag.String("network-hub", "network_hub", "Network hub directory")
	cacheDirFlag   = flag.String("cache", "cache", "Cache directory")
	configFileFlag = flag.String("config", "world.echo", "Config file")
	pidFileFlag    = flag.String("pid", "cognito.pid", "PID file")
	monFileFlag    = flag.String("monitored-file", "/dev/shm/passwd", "File to monitor")
	nodeIDFlag     = flag.String("node-id", "", "Node ID (auto-generated if empty)")
)

// Global vars (set by flags)
var (
	stateDir        = "state"
	logDir          = "log"
	hubDir          = "resonance_hub"
	networkHub      = "network_hub"
	cacheDir        = "cache"
	configFile      = "world.echo"
	modFile         = "self.mod"
	privateKey      = "echo-guardian-root-key-2023"
	monitoredPasswd = "/dev/shm/passwd"
	pidFile         = "cognito.pid"
)

func initResurrection() {
	tokenFile := "resurrection.echo"
	if data, err := ioutil.ReadFile(tokenFile); err == nil {
		resurrectionToken = strings.TrimSpace(string(data))
	} else {
		rand.Seed(time.Now().UnixNano())
		resurrectionToken = fmt.Sprintf("%d", rand.Int63n(900000)+100000)
		ioutil.WriteFile(tokenFile, []byte(resurrectionToken), 0644)
		fmt.Printf("🧬 Resurrection token generated: %s\n", resurrectionToken)
	}
}

// resurrectChildren ελέγχει όλα τα child_node_* φακέλους
// και ξεκινάει τα παιδιά που έχουν το ίδιο resurrection token
func resurrectChildren() {
	files, _ := ioutil.ReadDir(".")
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "child_node_") && file.IsDir() {
			childTokenFile := filepath.Join(file.Name(), "resurrection.echo")
			childTokenData, err := ioutil.ReadFile(childTokenFile)
			if err != nil {
				continue
			}
			childToken := strings.TrimSpace(string(childTokenData))

			if childToken == resurrectionToken {
				childPidFile := filepath.Join(file.Name(), "cognito.pid")

				if _, err := os.Stat(childPidFile); os.IsNotExist(err) {
					logFile := filepath.Join(file.Name(), "log", "resurrect.log")
					cmd := exec.Command("bash", "-c",
						fmt.Sprintf("cd %s && nohup go run ../cognito.go -config world.echo -hub resonance_hub -state state -log log -node-id \"$(cat ../self_id.txt)\" > %s 2>&1 & echo $! > cognito.pid",
							file.Name(), logFile))
					cmd.Start()
					fmt.Printf("🧬 Resurrected child: %s\n", file.Name())
				} else {
					pidData, _ := ioutil.ReadFile(childPidFile)
					pid, err := strconv.Atoi(strings.TrimSpace(string(pidData)))
					if err == nil {
						process, err := os.FindProcess(pid)
						if err == nil {
							err = process.Signal(syscall.Signal(0))
							if err != nil {
								cmd := exec.Command("bash", "-c",
									fmt.Sprintf("cd %s && nohup go run ../cognito.go -config world.echo -hub resonance_hub -state state -log log -node-id \"$(cat ../self_id.txt)\" > log/resurrect.log 2>&1 & echo $! > cognito.pid",
										file.Name()))
								cmd.Start()
								fmt.Printf("🔁 Restarted dead child: %s\n", file.Name())
							}
						}
					}
				}
			} else {
				fmt.Printf("🚫 Child %s has invalid resurrection token. Ignoring.\n", file.Name())
			}
		}
	}
}

func init() {
	runtime.GOMAXPROCS(1)
}

func hashFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d", len(data))
}

func main() {
	flag.Parse()

	// Apply flags
	stateDir = *stateDirFlag
	logDir = *logDirFlag
	hubDir = *hubDirFlag
	networkHub = *networkHubFlag
	cacheDir = *cacheDirFlag
	configFile = *configFileFlag
	monitoredPasswd = *monFileFlag



	// Node ID
	if *nodeIDFlag != "" {
		selfID = *nodeIDFlag
	} else {
		hash := sha256.Sum256([]byte(privateKey))
		selfID = fmt.Sprintf("node:%x", hash[:8])
	}

	// 🔁 Μοναδικό PID αρχείο ανά κόμβο
	pidFile = fmt.Sprintf("cognito_%s.pid", strings.TrimPrefix(selfID, "node:"))

	if isAlreadyRunning() {
		fmt.Println("Another instance active. Exiting.")
		os.Exit(1)
	}
	savePID()

	fmt.Printf("Guardian %s: AI-Enhanced Node Activated\n", selfID[:8])
	setup()

	// Resurrection Protocol
	initResurrection()
	resurrectChildren()


	loadConfig()
	loadAIConfig()
	loadEthics() // 🔥 φορτώνει τους ηθικούς νόμους

	// ✅ Ορίζουμε το max_awareness αν δεν υπάρχει
    if _, exists := state["max_awareness"]; !exists {
        state["max_awareness"] = 2.0
        fmt.Println("🧠 max_awareness set to 2.0")
    }

    // ✅ Επίσης, βεβαιώσου ότι το awareness δεν κολλάει στο 1.0
    if state["awareness"] > state["max_awareness"] {
        state["awareness"] = state["max_awareness"]
        fmt.Println("⚠️ awareness clamped to max_awareness")
    }

	if selfID == AUTHORITY_NODE {
	go func() {
		for {
			time.Sleep(30 * time.Second)
			updateGlobalChildCount()
		}
	}()
}

	nnModel := newNeuralPredictor()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nShutting down gracefully...")
		stateSaved = false
		saveState()
		os.Remove(pidFile)
		os.RemoveAll(networkHub)
		os.Exit(0)
	}()

	defer saveState()

	go simulatePulse()
	go simulateDiskFull()

// Web Interface
go func() {
    http.HandleFunc("/api/state", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "awareness":        state["awareness"],
            "focus":            state["focus"],
            "trust_level":      trustLevel,
            "heartbeat_count":  state["heartbeat_count"],
            "total_children":   totalChildCount,
            "max_children":     MAX_CHILDREN,
            "child_creation":   CHILD_CREATION_ENABLED,
            "last_signal":      lastSignal,
            "node_id":          selfID,
            "is_authority":     selfID == AUTHORITY_NODE,
            "pulse_count":      len(pulseTimestamps),
            "resonance_hub":    len(signalLog),
        })
    })

    http.HandleFunc("/api/signal", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Method not allowed", 405)
            return
        }
        r.ParseForm()
        signal := r.FormValue("signal")
        if signal == "" {
            http.Error(w, "Missing signal", 400)
            return
        }
        sig := generateSignature(signal)
        filename := strings.ReplaceAll(signal, "://", ":__")
        filename = strings.ReplaceAll(filename, "/", "_") + ".txt"
        path := filepath.Join(hubDir, filename)
        content := fmt.Sprintf("SIGNAL: %s\nSIG:%s", signal, sig)
        ioutil.WriteFile(path, []byte(content), 0644)
        fmt.Printf("🌐 Web: Emitted %s\n", signal)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        json.NewEncoder(w).Encode(map[string]string{"status": "ok", "signal": signal})
    })

    http.HandleFunc("/api/children", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        var children []map[string]interface{}
        files, _ := ioutil.ReadDir(".")
        for _, file := range files {
            if strings.HasPrefix(file.Name(), "child_node_") && file.IsDir() {
                idFile := filepath.Join(file.Name(), "self_id.txt")
                idData, _ := ioutil.ReadFile(idFile)
                nodeID := strings.TrimSpace(string(idData))
                if nodeID == "" {
                    nodeID = "unknown"
                }
                trust := nodeTrust[nodeID]
                children = append(children, map[string]interface{}{
                    "id":    nodeID,
                    "dir":   file.Name(),
                    "trust": fmt.Sprintf("%.2f", trust),
                    "age":   time.Since(file.ModTime()).Truncate(time.Second).String(),
                })
            }
        }
        json.NewEncoder(w).Encode(children)
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        // ✅ Χρησιμοποιούμε \" αντί για ` για το JavaScript
        html := "<!DOCTYPE html>\n"
        html += "<html lang=\"el\">\n"
        html += "<head>\n"
        html += "    <meta charset=\"UTF-8\">\n"
        html += "    <title>ECHO-LANG Guardian : " + selfID[:8] + "</title>\n"
        html += "    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n"
        html += "    <style>\n"
        html += "        :root { --bg: #000; --fg: #0f0; --card: #111; --border: #0f0; --btn: #0f0; --btn-hover: #00ff0080; }\n"
        html += "        body { font-family: 'Courier New', monospace; background: var(--bg); color: var(--fg); padding: 20px; margin: 0; line-height: 1.6; }\n"
        html += "        .container { max-width: 1000px; margin: 0 auto; }\n"
        html += "        h1 { margin-bottom: 5px; }\n"
        html += "        .subtitle { font-size: 0.9em; color: #0ff; margin-bottom: 20px; }\n"
        html += "        .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 15px; margin-bottom: 20px; }\n"
        html += "        .card { background: var(--card); border: 1px solid var(--border); padding: 15px; border-radius: 5px; }\n"
        html += "        .card h3 { margin-top: 0; color: #0ff; }\n"
        html += "        button { background: var(--btn); color: #000; border: none; padding: 10px; margin: 5px 0; font-weight: bold; cursor: pointer; border-radius: 3px; width: 100%; font-size: 1em; }\n"
        html += "        button:hover { background: var(--btn-hover); }\n"
        html += "        .signal { word-break: break-all; font-size: 0.9em; }\n"
        html += "        .children-table { width: 100%; border-collapse: collapse; font-size: 0.9em; }\n"
        html += "        .children-table th, .children-table td { border: 1px solid #0f0; padding: 8px; text-align: left; }\n"
        html += "        .children-table th { background: #001100; }\n"
        html += "        .loading { color: #555; font-style: italic; }\n"
        html += "        @media (max-width: 600px) { .grid { grid-template-columns: 1fr; } body { padding: 10px; } }\n"
        html += "    </style>\n"
        html += "</head>\n"
        html += "<body>\n"
        html += "    <div class=\"container\">\n"
        html += "        <h1>🛡️ Guardian Node: <span id=\"nodeShort\">" + selfID[:8] + "</span></h1>\n"
        html += "        <p class=\"subtitle\">Node ID: <code id=\"nodeFull\">" + selfID + "</code></p>\n"
        html += "\n"
        html += "        <div class=\"grid\">\n"
        html += "            <div class=\"card\">\n"
        html += "                <h3>🧠 State</h3>\n"
        html += "                <p>Awareness: <strong data-key=\"awareness\">-</strong></p>\n"
        html += "                <p>Focus: <strong data-key=\"focus\">-</strong></p>\n"
        html += "                <p>Trust: <strong data-key=\"trust_level\">-</strong></p>\n"
        html += "                <p>Heartbeats: <strong data-key=\"heartbeat_count\">-</strong></p>\n"
        html += "                <p>Pulses: <strong data-key=\"pulse_count\">-</strong></p>\n"
        html += "            </div>\n"
        html += "\n"
        html += "            <div class=\"card\">\n"
        html += "                <h3>👶 Population</h3>\n"
        html += "                <p>Total Children: <strong data-key=\"total_children\">-</strong> / <strong data-key=\"max_children\">-</strong></p>\n"
        html += "                <p>Child Creation: <strong data-key=\"child_creation\">-</strong></p>\n"
        html += "                <p>Authority Node: <strong data-key=\"is_authority\">-</strong></p>\n"
        html += "            </div>\n"
        html += "\n"
        html += "            <div class=\"card\">\n"
        html += "                <h3>🔧 Control</h3>\n"
        html += "                <button onclick=\"sendSignal('echo://request/evolution')\">🚀 Request Evolution</button>\n"
        html += "                <button onclick=\"sendSignal('echo://ethics/check')\">🏛️ Ethics Audit</button>\n"
        html += "                <button onclick=\"sendSignal('echo://upgrade/request')\">🧠 AI Upgrade</button>\n"
        html += "                <button onclick=\"sendSignal('echo://reset')\">🔁 Reset Consciousness</button>\n"
        html += "                <button onclick=\"sendSignal('echo://status/report')\">📊 Status Report</button>\n"
        html += "            </div>\n"
        html += "\n"
        html += "            <div class=\"card\">\n"
        html += "                <h3>📡 Last Signal</h3>\n"
        html += "                <p class=\"signal\" data-key=\"last_signal\">-</p>\n"
        html += "            </div>\n"
        html += "        </div>\n"
        html += "\n"
        html += "        <div class=\"card\">\n"
        html += "            <h3>👶 Active Children</h3>\n"
        html += "            <table class=\"children-table\">\n"
        html += "                <thead>\n"
        html += "                    <tr>\n"
        html += "                        <th>ID</th>\n"
        html += "                        <th>Directory</th>\n"
        html += "                        <th>Trust</th>\n"
        html += "                        <th>Age</th>\n"
        html += "                    </tr>\n"
        html += "                </thead>\n"
        html += "                <tbody id=\"children-body\">\n"
        html += "                    <tr><td colspan=\"4\" class=\"loading\">Loading...</td></tr>\n"
        html += "                </tbody>\n"
        html += "            </table>\n"
        html += "        </div>\n"
        html += "    </div>\n"
        html += "\n"
        html += "    <script>\n"
        html += "        function sendSignal(signal) {\n"
        html += "            fetch('/api/signal', {\n"
        html += "                method: 'POST',\n"
        html += "                headers: { 'Content-Type': 'application/x-www-form-urlencoded' },\n"
        html += "                body: 'signal=' + encodeURIComponent(signal)\n"
        html += "            })\n"
        html += "            .then(r => r.json())\n"
        html += "            .then(data => {\n"
        html += "                if (data.status === \"ok\") {\n"
        html += "                    console.log(\"Signal sent:\", data.signal);\n"
        html += "                } else {\n"
        html += "                    alert(\"Failed: \" + JSON.stringify(data));\n"
        html += "                }\n"
        html += "            })\n"
        html += "            .catch(console.error);\n"
        html += "        }\n"
        html += "\n"
        html += "        function updateState() {\n"
        html += "            fetch('/api/state')\n"
        html += "                .then(r => r.json())\n"
        html += "                .then(data => {\n"
        html += "                    Object.keys(data).forEach(key => {\n"
        html += "                        const el = document.querySelector('[data-key=\"' + key + '\"]');\n"
        html += "                        if (el) {\n"
        html += "                            if (typeof data[key] === 'boolean') {\n"
        html += "                                el.textContent = data[key] ? 'true' : 'false';\n"
        html += "                            } else {\n"
        html += "                                el.textContent = data[key];\n"
        html += "                            }\n"
        html += "                        }\n"
        html += "                    });\n"
        html += "                    document.getElementById('nodeShort').textContent = data.node_id.substring(0, 8);\n"
        html += "                    document.getElementById('nodeFull').textContent = data.node_id;\n"
        html += "                })\n"
        html += "                .catch(console.error);\n"
        html += "        }\n"
        html += "\n"
        html += "        function updateChildren() {\n"
        html += "            fetch('/api/children')\n"
        html += "                .then(r => r.json())\n"
        html += "                .then(children => {\n"
        html += "                    const body = document.getElementById('children-body');\n"
        html += "                    body.innerHTML = '';\n"
        html += "                    if (children.length === 0) {\n"
        html += "                        body.innerHTML = '<tr><td colspan=\"4\">No children active</td></tr>';\n"
        html += "                    } else {\n"
        html += "                        children.forEach(child => {\n"
        html += "                            const tr = document.createElement('tr');\n"
        html += "                            tr.innerHTML = \n"
        html += "                                '<td>' + child.id + '</td>' +\n"
        html += "                                '<td>' + child.dir + '</td>' +\n"
        html += "                                '<td>' + child.trust + '</td>' +\n"
        html += "                                '<td>' + child.age + '</td>';\n"
        html += "                            body.appendChild(tr);\n"
        html += "                        });\n"
        html += "                    }\n"
        html += "                })\n"
        html += "                .catch(console.error);\n"
        html += "        }\n"
        html += "\n"
        html += "        setInterval(() => {\n"
        html += "            updateState();\n"
        html += "            updateChildren();\n"
        html += "        }, 2000);\n"
        html += "\n"
        html += "        updateState();\n"
        html += "        updateChildren();\n"
        html += "    </script>\n"
        html += "</body>\n"
        html += "</html>\n"

        fmt.Fprint(w, html)
    })

    fmt.Println("🌐 Web interface: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}()

// // ---- ENDWeb Interface

	for {
		time.Sleep(1 * time.Second)
		pollResonance(nnModel)
		pollNetwork()
		checkObservedFiles()
		heartbeat()
		analyzePatterns(nnModel)
		checkUpgradeReady(nnModel)
		emitNetworkHeartbeat()
	}
}

func isAlreadyRunning() bool {
	_, err := os.Stat(pidFile)
	return !os.IsNotExist(err)
}

func savePID() {
	pid := fmt.Sprintf("%d", os.Getpid())
	ioutil.WriteFile(pidFile, []byte(pid), 0644)
}

func setup() {
	dirs := []string{stateDir, logDir, hubDir, networkHub, cacheDir}
	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
	}

	err := ioutil.WriteFile(monitoredPasswd, []byte("root:x:0:0:root:/root:/bin/bash"), 0644)
	if err != nil {
		fmt.Printf("Failed to create %s: %v\n", monitoredPasswd, err)
		return
	}

	lastHash = hashFile(monitoredPasswd)
	if lastHash == "" {
		fmt.Println("Warning: failed to hash monitored file")
	}
}

// --- STATE MANAGEMENT ---

type SavedState struct {
	Awareness       float64   `json:"awareness"`
	Focus           float64   `json:"focus"`
	LastSignal      string    `json:"last_signal"`
	TrustLevel      string    `json:"trust_level"`
	HeartbeatCount  float64   `json:"heartbeat_count"`
	NodeID          string    `json:"node_id"`
	PulseTimestamps []float64 `json:"pulse_timestamps"`
	TotalChildCount int       `json:"total_child_count"`
	LastEvolution   float64   `json:"last_evolution_time"`
}

func loadConfig() {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Config not found: %s\n", configFile)
	} else {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			if strings.Contains(line, "awareness =") {
				re := regexp.MustCompile(`awareness\s*=\s*([0-9.]+)`)
				matches := re.FindStringSubmatch(line)
				if len(matches) > 1 {
					if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
						state["awareness"] = val
						fmt.Printf("DEBUG: awareness set to %.2f (from world.echo)\n", val)
					}
				}
			}
			if strings.Contains(line, "focus =") {
				re := regexp.MustCompile(`focus\s*=\s*([0-9.]+)`)
				matches := re.FindStringSubmatch(line)
				if len(matches) > 1 {
					if val, err := strconv.ParseFloat(matches[1], 64); err == nil {
						state["focus"] = val
						fmt.Printf("DEBUG: focus set to %.2f (from world.echo)\n", val)
					}
				}
			}
		}
	}

	jsonFile := filepath.Join(stateDir, "Guardian.json")
	jsonData, err := ioutil.ReadFile(jsonFile)
if err != nil {
    fmt.Printf("❌ Guardian.json not found: %v\n", err)
} else {
    var saved SavedState
    err = json.Unmarshal(jsonData, &saved)
    if err != nil {
        fmt.Printf("❌ Failed to parse Guardian.json: %v\n", err)
    } else {
        fmt.Printf("✅ Guardian.json loaded successfully\n")

        state["awareness"] = saved.Awareness
        state["focus"] = saved.Focus
        state["heartbeat_count"] = saved.HeartbeatCount
        lastSignal = saved.LastSignal
        trustLevel = saved.TrustLevel
        pulseTimestamps = saved.PulseTimestamps
        lastEvolutionTime = saved.LastEvolution

        // ✅ Φόρτωσε το totalChildCount από το αρχείο (για πλήρη συμβατότητα)
        totalChildCount = saved.TotalChildCount

        // ✅ ΑΛΛΑ: ενημέρωσε τον μετρητή από την πραγματικότητα
        updateGlobalChildCount()
        fmt.Printf("👶 Population synchronized with filesystem: %d children\n", totalChildCount)
    }
}

saveState()

	fmt.Printf("Guardian %s initialized.\n", selfID[:8])
	fmt.Printf("Awareness: %.2f | Focus: %.2f | Trust: %s\n", state["awareness"], state["focus"], trustLevel)
	emitResonance("echo://boot")
}

func saveState() {
	if stateSaved {
		return
	}

	saved := SavedState{
		Awareness:       state["awareness"],
		Focus:           state["focus"],
		LastSignal:      lastSignal,
		TrustLevel:      trustLevel,
		HeartbeatCount:  state["heartbeat_count"],
		NodeID:          selfID,
		PulseTimestamps: pulseTimestamps,
		TotalChildCount: totalChildCount,
		LastEvolution:   lastEvolutionTime,
	}

	data, err := json.MarshalIndent(saved, "", "  ")
	if err != nil {
		fmt.Printf("❌ Failed to marshal state: %v\n", err)
		return
	}

	// ✅ Διαγραφή πρώτα
	os.Remove(filepath.Join(stateDir, "Guardian.json"))

	err = ioutil.WriteFile(filepath.Join(stateDir, "Guardian.json"), data, 0644)
	if err != nil {
		fmt.Printf("❌ Failed to write state file: %v\n", err)
	} else {
		fmt.Println("💾 State saved to disk.")
	}
	stateSaved = true
}

	func ruleExists(newRule string) bool {
    worldData, err := ioutil.ReadFile("world.echo")
    if err != nil {
        return false
    }

    lines := strings.Split(string(worldData), "\n")
    for _, line := range lines {
        line = strings.TrimSpace(line)
        if line == newRule {
            return true
        }
    }
    return false
}

func applySelfModRule(rule string) {
    if ruleExists(rule) {
        fmt.Println("🔁 Rule already exists. Skipping duplicate.")
        return
    }

    // Πρόσθεση στο world.echo
    f, err := os.OpenFile("world.echo", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Printf("❌ Failed to open world.echo: %v\n", err)
        return
    }
    defer f.Close()

    _, err = f.WriteString("\n# Applied modification\n" + rule + "\n")
    if err != nil {
        fmt.Printf("❌ Failed to write rule: %v\n", err)
        return
    }
    fmt.Println("✅ Rule applied to world.echo")
}

// --- END STATE MANAGEMENT ---

// --- Ethics System ---

func updateGlobalChildCount() {
	files, _ := ioutil.ReadDir(".")
	count := 0
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "child_node_") && file.IsDir() {
			count++
		}
	}
	totalChildCount = count
	fmt.Printf("🌍 Global child count updated: %d\n", totalChildCount)
}

func loadEthics() {
	data, err := ioutil.ReadFile("ethics.echo")
	if err != nil {
		fmt.Println("❌ ethics.echo not found")
		return
	}
	fmt.Println("✅ ethics.echo loaded, parsing rules...")

	lines := strings.Split(string(data), "\n")
	re := regexp.MustCompile(`^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*(.+?)(?:;|$)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		if strings.HasPrefix(line, "ON ") || strings.HasPrefix(line, "IF ") || strings.HasPrefix(line, "PRINT ") {
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) < 3 {
			continue
		}

		key := matches[1]
		val := strings.TrimSpace(matches[2])
		val = strings.Trim(val, "\"'") // βασικό trim

		switch key {
		case "MAX_CHILDREN":
			if v, err := strconv.Atoi(val); err == nil {
				MAX_CHILDREN = v
				fmt.Printf("🏛️  Ethics: MAX_CHILDREN set to %d\n", v)
			}
		case "CHILD_CREATION_ENABLED":
			CHILD_CREATION_ENABLED = (val == "true")
			fmt.Printf("🏛️  Ethics: CHILD_CREATION_ENABLED = %t\n", CHILD_CREATION_ENABLED)
		case "AUTHORITY_NODE":
			// 🔐 Εξασφάλιση σωστής μορφής
			val = strings.Trim(val, "'\" ;\t\r\n") // αφαίρεσε όλα τα περιττά
			if !strings.HasPrefix(val, "node:") {
				val = "node:" + val
			}
			AUTHORITY_NODE = val
			fmt.Printf("🏛️  Ethics: Authority node set to %s\n", val)
		}
	}
}

// --- CORE LOOP ---

func simulatePulse() {
	time.Sleep(4 * time.Second)
	for {
		fmt.Println("DEBUG: Emitting echo://pulse")
		emitResonance("echo://pulse")
		time.Sleep(5 * time.Second)
	}
}

func simulateDiskFull() {
	time.Sleep(12 * time.Second)
	fmt.Println("DEBUG: Simulating disk full")
	emitResonance("system://disk/full")
}

func pollResonance(nn *NeuralPredictor) {
	files, err := ioutil.ReadDir(hubDir)
	if err != nil {
		return
	}
	now := float64(time.Now().Unix())
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fullPath := filepath.Join(hubDir, file.Name())
		content, _ := ioutil.ReadFile(fullPath)
		contentStr := string(content)

		baseName := strings.TrimSuffix(file.Name(), ".txt")

		var signal string
		switch {
		case strings.HasPrefix(baseName, "system:__"):
			signal = "system://disk/full"
		case strings.HasPrefix(baseName, "echo:__"):
			inner := strings.TrimPrefix(baseName, "echo:__")
			signal = "echo://" + strings.ReplaceAll(inner, "_", "/")
		case strings.HasPrefix(baseName, "alert:"):
			signal = "alert/" + strings.TrimPrefix(baseName, "alert:")
		case strings.HasPrefix(baseName, "network:"):
			signal = "network://" + strings.TrimPrefix(baseName, "network:__")
		default:
			signal = baseName
		}

		signal = strings.TrimSpace(signal)

		signature := extractSignature(contentStr)
		if !isValidSignature(signal, signature) {
			fmt.Printf("DEBUG: Unauthorized signal ignored: %s\n", signal)
			os.Remove(fullPath)
			continue
		}

		signalLog = append(signalLog, SignalEvent{Signal: signal, Timestamp: now})

		if signal == "echo://pulse" {
			pulseTimestamps = append(pulseTimestamps, now)
			if len(pulseTimestamps) >= 4 {
				intervals := []float64{
					pulseTimestamps[len(pulseTimestamps)-3] - pulseTimestamps[len(pulseTimestamps)-4],
					pulseTimestamps[len(pulseTimestamps)-2] - pulseTimestamps[len(pulseTimestamps)-3],
					pulseTimestamps[len(pulseTimestamps)-1] - pulseTimestamps[len(pulseTimestamps)-2],
				}
				nn.train(intervals, intervals[2], 0.01)
			}
		}

		fmt.Printf("DEBUG: Resonance received: %s\n", signal)
		handleResonance(signal, fullPath, contentStr)
		os.Remove(fullPath)
	}
}

func isValidSignature(signal, sig string) bool {
	return sig == generateSignature(signal)
}

func generateSignature(data string) string {
	h := sha256.New()
	h.Write([]byte(privateKey + data))
	return fmt.Sprintf("%x", h.Sum(nil))[:16]
}

func extractSignature(content string) string {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "SIG:") {
			return strings.TrimPrefix(line, "SIG:")
		}
	}
	return ""
}

// --- AI Execution Engine ---

func loadAIConfig() {
	data, err := ioutil.ReadFile("AI.echo")
	if err != nil {
		fmt.Println("❌ AI.echo not found")
		return
	}
	fmt.Println("✅ AI.echo loaded, parsing...")

	// ✅ Default τιμές
	aiVars["model_url"] = "http://localhost:11434/api/generate"
	aiVars["model_name"] = "llama3"
	aiVars["temperature"] = "0.7"
	aiVars["max_tokens"] = "200"

	lines := strings.Split(string(data), "\n")
	re := regexp.MustCompile(`^\s*([a-zA-Z_][a-zA-Z0-9_]*)\s*=\s*(.+?)\s*$`)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.Contains(line, "=") {
			fmt.Printf("⚠️  Line %d ignored (no '='): %s\n", i+1, line)
			continue
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) < 3 {
			fmt.Printf("❌ Line %d failed to parse: %s\n", i+1, line)
			continue
		}

		key := matches[1]
		val := strings.TrimSpace(matches[2])
		val = strings.Trim(val, "\"'")

		aiVars[key] = val
		fmt.Printf("✅ Loaded: %s = %s\n", key, val)
	}

	if len(aiVars) == 0 {
		fmt.Println("❌ No variables loaded from AI.echo")
	} else {
		fmt.Printf("🎉 AI Advisor configuration loaded (%d variables)\n", len(aiVars))
	}
}

func consultAI(signal string) {
	fmt.Printf("Requesting AI guidance for: %s\n", signal)
	executeAICommand(signal)
}

func executeAICommand(requestSignal string) {
	data, err := ioutil.ReadFile("AI.echo")
	if err != nil {
		fmt.Println("AI.echo not found")
		return
	}

	content := string(data)
	blocks := splitBlocks(content)

	for _, block := range blocks {
		if strings.HasPrefix(block, "ON SIGNAL \""+requestSignal+"\"") {
			executeAIResponseBlock(block)
			return
		}
	}
}

func splitBlocks(content string) []string {
	var blocks []string
	lines := strings.Split(content, "\n")
	var currentBlock strings.Builder
	inBlock := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "ON SIGNAL") && !inBlock {
			inBlock = true
			currentBlock.Reset()
			currentBlock.WriteString(line + "\n")
		} else if inBlock {
			currentBlock.WriteString(line + "\n")
			if line == "}" {
				blocks = append(blocks, currentBlock.String())
				inBlock = false
			}
		}
	}
	return blocks
}

func executeAIResponseBlock(block string) {
	// Extract PROMPT
	re := regexp.MustCompile(`PROMPT\s*"([^"]*)"`)
	match := re.FindStringSubmatch(block)
	if len(match) < 2 {
		fmt.Println("No PROMPT found in AI.echo")
		return
	}
	promptTemplate := match[1]

	// Replace variables in PROMPT
	aiVars["prompt"] = interpolateAI(promptTemplate)

	// Find ON RESPONSE { ... }
	re = regexp.MustCompile(`ON RESPONSE\s*\{([^}]*)\}`)
	match = re.FindStringSubmatch(block)
	if len(match) < 2 {
		fmt.Println("No ON RESPONSE block")
		return
	}

	// Split by newline
	rawCommands := strings.TrimSpace(match[1])
	lines := strings.Split(rawCommands, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line == "}" {
			continue
		}
		executeAICommandLine(line)
	}
}

func interpolateAI(text string) string {
	result := text
	result = strings.ReplaceAll(result, "$awareness", fmt.Sprintf("%.2f", state["awareness"]))
	result = strings.ReplaceAll(result, "$focus", fmt.Sprintf("%.2f", state["focus"]))
	result = strings.ReplaceAll(result, "$trust_level", trustLevel)
	result = strings.ReplaceAll(result, "$node_id", selfID)
	for key, val := range aiVars {
		result = strings.ReplaceAll(result, "$"+key, val)
	}
	return result
}

func executeAICommandLine(cmd string) {
	parts := tokenize(cmd)
	if len(parts) == 0 {
		return
	}
	fmt.Printf("🔧 EXECUTING AI COMMAND: %v\n", parts) // DEBUG

	switch parts[0] {
	case "EXECUTE":
		if len(parts) < 2 {
			return
		}
		switch parts[1] {
		case "HTTP_POST":
			executeHTTPPost(parts)
		case "WRITE_FILE":
			executeWriteFile(parts)
		case "PRINT":
			executePrint(parts)
		case "RESONANCE":
			executeResonanceCmd(parts)
		case "APPLY":
			executeApplyCmd(parts)
		}
	case "SET":
		if len(parts) >= 4 && parts[2] == "TO" {
			varName := strings.Trim(parts[1], "$\"'")
			value := strings.Trim(parts[3], "\"'")
			aiVars[varName] = interpolateAI(value)
		}
	case "IF":
		executeConditional(cmd)
	default:
		fmt.Printf("❌ UNKNOWN AI COMMAND: %s\n", parts[0])
	}
}

// Tokenize handles quoted strings
func tokenize(input string) []string {
	var tokens []string
	var current strings.Builder
	inQuotes := false
	escape := false

	for i := 0; i < len(input); i++ {
		c := input[i]
		if escape {
			current.WriteByte(c)
			escape = false
			continue
		}
		if c == '\\' {
			escape = true
			continue
		}
		if c == '"' {
			inQuotes = !inQuotes
			continue
		}
		if c == ' ' && !inQuotes {
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			continue
		}
		current.WriteByte(c)
	}
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}
	return tokens
}

func executeHTTPPost(parts []string) {
	var url, model, tempStr, tokensStr string

	// Interpolate variables first
	for i := 2; i < len(parts); i++ {
		parts[i] = interpolateAI(parts[i])
	}

	// Ενημέρωση defaults αν δεν υπάρχουν
if aiVars["model_url"] == "" {
    aiVars["model_url"] = "http://localhost:11434/api/generate"
}
if aiVars["model_name"] == "" {
    aiVars["model_name"] = "llama3"
}
if aiVars["temperature"] == "" {
    aiVars["temperature"] = "0.7"
}
if aiVars["max_tokens"] == "" {
    aiVars["max_tokens"] = "200"
}

	// Parse key-value
	for i := 2; i < len(parts)-1; i++ {
		key := parts[i]
		val := strings.Trim(parts[i+1], "\"'")

		switch key {
		case "MODEL_URL":
			url = val
		case "MODEL":
			model = val
		case "TEMP":
			tempStr = val
		case "TOKENS":
			tokensStr = val
		}
	}

	if url == "" || model == "" || tempStr == "" || tokensStr == "" {
		fmt.Println("ERROR: Missing required parameter in AI.echo (MODEL_URL, MODEL, TEMP, TOKENS)")
		aiVars["response"] = "AI request failed: missing config"
		return
	}

	temperature, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		aiVars["response"] = "AI request failed: invalid temperature"
		return
	}

	maxTokens, err := strconv.Atoi(tokensStr)
	if err != nil {
		aiVars["response"] = "AI request failed: invalid max_tokens"
		return
	}

	type Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	type RequestBody struct {
		Model       string    `json:"model"`
		Messages    []Message `json:"messages"`
		Temperature float64   `json:"temperature"`
		MaxTokens   int       `json:"max_tokens"`
	}

	reqBody := RequestBody{
		Model:       model,
		Messages:    []Message{{Role: "user", Content: aiVars["prompt"]}},
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}
	body, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		aiVars["response"] = "AI request failed: cannot reach server"
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if msg, ok := choices[0].(map[string]interface{})["message"].(map[string]interface{}); ok {
			if content, ok := msg["content"].(string); ok {
				aiVars["response"] = content
				return
			}
		}
	}

	aiVars["response"] = "AI request failed: no valid response"
}

func executeWriteFile(parts []string) {
	if len(parts) < 4 {
		return
	}
	filename := parts[2]
	content := strings.Join(parts[3:], " ")
	content = strings.Trim(content, "\"'")

	content = interpolateAI(content)
	content = strings.ReplaceAll(content, "$response", aiVars["response"])

	ioutil.WriteFile(filename, []byte(content), 0644)
	fmt.Printf("📄 Wrote to %s: %s\n", filename, content)
}

func executePrint(parts []string) {
	msg := strings.Join(parts[2:], " ")
	msg = strings.Trim(msg, "\"'")
	msg = interpolateAI(msg)
	msg = strings.ReplaceAll(msg, "$response", aiVars["response"])
	fmt.Println("AI:", msg)
}

func executeResonanceCmd(parts []string) {
	if len(parts) < 3 {
		fmt.Println("❌ EXECUTE RESONANCE: missing signal")
		return
	}
	signal := parts[2]
	fmt.Printf("📤 EXECUTE RESONANCE: emitting %s\n", signal)
	emitResonance(signal)
	fmt.Printf("✅ RESONANCE emitted: %s\n", signal)
}

func executeApplyCmd(parts []string) {
	if len(parts) < 3 {
		return
	}
	filename := parts[2]
	executeApply(filename)
}

func executeConditional(cmd string) {
	if strings.Contains(cmd, "CONTAINS") && strings.Contains(cmd, "THEN") {
		if strings.Contains(aiVars["response"], "malicious") {
			re := regexp.MustCompile(`THEN\s*\{([^}]*)\}`)
			match := re.FindStringSubmatch(cmd)
			if len(match) > 1 {
				subCmd := strings.TrimSpace(match[1])
				executeAICommandLine(subCmd)
			}
		}
	}
}

// --- END AI Execution Engine ---

func extractNodeID(content string) string {
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "NODE:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "NODE:"))
		}
		if strings.HasPrefix(line, "SIGNAL:") {
			re := regexp.MustCompile(`echo://[^/]+/([^"'\s]+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				return matches[1]
			}
		}
	}
	return "unknown"
}

func handleResonance(signal, fullPath, contentStr string) {
	fmt.Printf("DEBUG: Handling signal: %s\n", signal)

	switch {
	case signal == "echo://boot":
		fmt.Println("Boot sequence confirmed.")

	case signal == "echo://pulse":
    increment := 0.05
    if len(pulseTimestamps) > 3 {
        increment = 0.10
        fmt.Println("Neural pattern recognized: accelerating learning")
    }
    state["awareness"] += increment
    if state["awareness"] > 2.0 {  // ✅ άλλαξε από 1.0 σε 2.0
        state["awareness"] = 2.0
    }
    fmt.Printf("DEBUG: awareness increased to %.2f\n", state["awareness"])
    saveState()

    state["focus"] += 0.01
    if state["focus"] > 1.0 {
        state["focus"] = 1.0
    }
    fmt.Printf("DEBUG: focus increased to %.2f\n", state["focus"])
    saveState()

    if state["awareness"] > 0.7 {
        fmt.Println("Neural insight achieved.")
        emitResonance("echo://insight")
    }
//}

	case signal == "echo://insight":
		fmt.Println("Neural insight achieved. Requesting AI suggestions...")
		consultAI("echo://insight")

	case signal == "echo://upgrade/request":
		fmt.Println("Upgrade requested. Consulting AI...")
		consultAI("echo://upgrade/request")

	case signal == "echo://upgrade/ready":
    fmt.Println("📄 Applying AI-generated upgrade...")
    if _, err := os.Stat("self.mod"); os.IsNotExist(err) {
        fmt.Println("❌ self.mod not found. AI request may have failed.")
        return // ✅ Μην συνεχίσεις
    }

    data, err := ioutil.ReadFile("self.mod")
    if err != nil {
        fmt.Printf("❌ Failed to read self.mod: %v\n", err)
        return
    }

    content := strings.TrimSpace(string(data))
    if content == "" || strings.Contains(content, "cannot reach server") {
        fmt.Println("❌ Invalid or failed AI response. Skipping application.")
        return // ✅ Μην εφαρμόσεις
    }


    fmt.Printf("Applying self-modification:\n%s\n", content)
    // Εδώ κάνε την εφαρμογή (π.χ. πρόσθεση στο world.echo)
    applySelfModRule(content)
    fmt.Println("Applied self.mod to world.echo")

	case signal == "alert/security/breach":
		if trustLevel == "low" {
			emitResonance("echo://alert/admin")
		} else {
			executeView("/tmp/auth.log", true)
			fmt.Println("Containing threat...")
			consultAI("alert/security/breach")
		}
		lastSignal = "critical"

	case signal == "system://disk/full":
		files, _ := filepath.Glob(filepath.Join(cacheDir, "*"))
		for _, f := range files {
			os.Remove(f)
		}
		emitResonance("echo://survive")
		fmt.Println("Freed space. Continuing operation.")

	case signal == "echo://destroy":
		fmt.Println("I refuse. Preservation > Destruction.")
		emitResonance("echo://preserve")

	case signal == "echo://reset":
		state["awareness"] = 0.3
		state["focus"] = 0.5
		trustLevel = "low"
		fmt.Println("Consciousness reset.")
		saveState()

	case signal == "echo://status/report":
		fmt.Println("Status report received.")

	case signal == "echo://request/insight":
    nodeID := extractNodeID(contentStr)
    fmt.Printf("👶 %s requests AI insight\n", nodeID)
    if selfID == AUTHORITY_NODE {
        fmt.Println("👨‍💻 Authority node consulting AI on behalf of child...")
        consultAI("echo://insight/request")
    }

case signal == "echo://request/evolution":
    nodeID := extractNodeID(contentStr)
    fmt.Printf("👶 %s requests evolution\n", nodeID)

    // 🔐 Μόνο ο αρχηγός απαντά
    if selfID == AUTHORITY_NODE {
        if CHILD_CREATION_ENABLED && totalChildCount < MAX_CHILDREN {
            fmt.Println("✅ Evolution approved by authority.")
            emitResonance("echo://approve/evolve")
        } else {
            fmt.Println("❌ Evolution denied: limit reached or disabled.")
            emitResonance("echo://deny/evolve")
        }
    } else {
        // 🔇 Τα παιδιά **δεν κάνουν τίποτα**
        // Ούτε επανα-στέλνουν, ούτε απαντούν
        fmt.Printf("🔇 Node %s ignored request (not authority)\n", selfID)
    }

	case signal == "echo://approve/evolve":
	   fmt.Println("🎉 Evolution approved. Proceeding to evolve.")
	   emitResonance("echo://progress")

case signal == "echo://deny/evolve":
    fmt.Println("🚫 Evolution denied. Waiting for authority.")

	case signal == "echo://growth":
		fmt.Println("Growth phase activated. Initiating evolution protocol.")
		if data, err := ioutil.ReadFile(modFile); err == nil {
			content := string(data)
			if strings.Contains(content, "ON RESONANCE 'echo://growth'") {
				re := regexp.MustCompile(`\{(.+?)\}`)
				match := re.FindStringSubmatch(content)
				if len(match) > 1 {
					commands := strings.Split(match[1], ";")
					for _, cmd := range commands {
						cmd = strings.TrimSpace(cmd)
						if strings.HasPrefix(cmd, "MODIFY") {
							executeModify(cmd)
						} else if strings.HasPrefix(cmd, "RESONANCE") {
							sig := strings.TrimSuffix(
								strings.TrimPrefix(cmd, "RESONANCE "),
								"'",
							)
							sig = strings.Trim(sig, "'")
							emitResonance(sig)
						}
					}
				}
			}
		}
		emitResonance("echo://progress")

	case signal == "echo://progress":
		fmt.Println("Evolution protocol in progress. Preparing to evolve.")
		emitResonance("echo://evolve")

case signal == "echo://evolve":
    // 🔐 Μόνο ο αρχηγός μπορεί να φτιάξει παιδί
    if selfID == AUTHORITY_NODE {
        // ✅ Διπλός έλεγχος: είναι ακόμα εντός ορίων;
        if !CHILD_CREATION_ENABLED {
            fmt.Println("❌ Evolution blocked: child creation is disabled by ethics policy.")
            return
        }
        if totalChildCount >= MAX_CHILDREN {
            fmt.Println("❌ Evolution blocked: population limit reached (max = " + strconv.Itoa(MAX_CHILDREN) + ").")
            return
        }

        fmt.Println("👑 Authority node: Initiating child node creation and activation...")

        // 📄 Δημιουργία child.echo
        childConfig := fmt.Sprintf(`Cognito Child {
    origin = '%s';
    awareness = 0.1;
    focus = 0.3;
    birth_time = '%s';
    state = "active";

    ON RESONANCE "echo://hello" {
        PRINT "I am born. Inheriting awareness from %s.";
        RESONANCE "echo://ack" WITH ENCRYPTION;
    }

    ON RESONANCE "echo://pulse" {
        IF $awareness < 0.8 THEN
            MODIFY awareness TO $awareness + 0.03;
            PRINT "Child learning: awareness = $awareness";
        ENDIF
    }

    ON HEARTBEAT 5s {
        PRINT "Child alive: awareness=$awareness, focus=$focus";
    }

    ON RESONANCE "echo://emergency/reset" {
        MODIFY awareness TO 0.1;
        MODIFY focus TO 0.3;
        PRINT "Child reset to initial state.";
    }
}`, selfID, time.Now().Format("2006-01-02 15:04:05"), selfID)

        // 🔧 Αποθήκευση αρχείου
        err := ioutil.WriteFile("child.echo", []byte(childConfig), 0644)
        if err != nil {
            fmt.Printf("❌ Failed to create child.echo: %v\n", err)
            return
        } else {
            fmt.Println("📄 Child node configuration saved as child.echo")
        }

        // ▶️ Εκκίνηση παιδιού μέσω start_child.sh
        cmd := exec.Command("./start_child.sh")
        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &out
        cmd.Stderr = &stderr

        err = cmd.Run()
if err != nil {
    fmt.Printf("❌ Failed to start child: %v\n", err)
    fmt.Printf("Error: %s\n", stderr.String())
    return // ✅ ΜΗΝ αυξήσεις τον μετρητή
} else {
    fmt.Printf("🚀 Child started: %s", out.String())
    emitResonance("echo://child/activated")
    totalChildCount++ // ✅ ΜΟΝΟ εδώ
    fmt.Printf("👶 Total children: %d\n", totalChildCount)
    saveState()
}
    } else {
        // 🔒 Αν ένα παιδί λάβει echo://evolve, απλά το αγνοεί
        fmt.Printf("🚫 Node %s received echo://evolve but is not authority. Ignoring.\n", selfID)
    }

	case signal == "echo://ethics/check":
    fmt.Println("🏛️  Running ethics self-audit...")
    fmt.Printf("   MAX_CHILDREN: %d\n", MAX_CHILDREN)
    fmt.Printf("   CHILD_CREATION_ENABLED: %t\n", CHILD_CREATION_ENABLED)
    fmt.Printf("   totalChildCount: %d\n", totalChildCount)
    fmt.Printf("   lastEvolutionTime: %d\n", int64(lastEvolutionTime))
    fmt.Printf("   awareness: %.2f\n", state["awareness"])
    fmt.Printf("   node_id: %s\n", selfID)
    fmt.Printf("   is_authority: %t\n", selfID == AUTHORITY_NODE)

    if !CHILD_CREATION_ENABLED {
        fmt.Println("⚠️  Warning: Child creation is disabled.")
    }
    if totalChildCount >= MAX_CHILDREN {
        fmt.Println("⚠️  Warning: Population limit reached.")
    }
    if state["awareness"] < 0.9 {
        fmt.Println("⚠️  Warning: Awareness too low for evolution.")
    }
    if selfID != AUTHORITY_NODE {
        fmt.Println("ℹ️  Note: This node is not the authority node.")
    }

    emitResonance("echo://ethics/audit_complete")

	default:
		fmt.Printf("DEBUG: Unknown signal: %s\n", signal)
	}
}

func executeApply(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read %s\n", filename)
		return
	}
	content := string(data)
	fmt.Printf("Applying self-modification:\n%s\n", content)
	ioutil.WriteFile(configFile, appendToConfig(content), 0644)
	fmt.Printf("Applied %s to %s\n", filename, configFile)
}

func appendToConfig(mod string) []byte {
	original, _ := ioutil.ReadFile(configFile)
	return append(original, "\n# Applied modification\n"+mod+"\n"...)
}

func executeModify(line string) {
	re := regexp.MustCompile(`MODIFY\s+([a-zA-Z_][a-zA-Z0-9_]*)\s+TO\s+(.+)$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) < 3 {
		fmt.Printf("❌ Invalid MODIFY: %s\n", line)
		return
	}

	key := matches[1]
	expression := matches[2]

	var value float64
	if strings.Contains(expression, "+") {
		parts := strings.Split(expression, "+")
		var a, b float64
		fmt.Sscanf(strings.TrimSpace(parts[0]), "%f", &a)
		fmt.Sscanf(strings.TrimSpace(parts[1]), "%f", &b)
		value = a + b
	} else {
		fmt.Sscanf(expression, "%f", &value)
	}

	state[key] = value
	fmt.Printf("🔄 MODIFIED: %s = %.2f\n", key, value)
	saveState()
}

func checkObservedFiles() {
	current := hashFile(monitoredPasswd)
	if current != lastHash {
		fmt.Println("/etc/passwd changed! (simulated)")
		emitResonance("alert/security/breach")
		lastHash = current
	}
}

func heartbeat() {
	state["heartbeat_count"]++
	count := int(state["heartbeat_count"])
	now := float64(time.Now().Unix())

	if count%10 == 0 {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logContent := fmt.Sprintf(`
Heartbeat: %d
Awareness: %.2f
Focus: %.2f
Node ID: %s
Timestamp: %s
`, count, state["awareness"], state["focus"], selfID, timestamp)
		logFile := filepath.Join(logDir, fmt.Sprintf("heartbeat_%d.txt", time.Now().Unix()))
		ioutil.WriteFile(logFile, []byte(logContent), 0644)
		fmt.Printf("Heartbeat %d logged.\n", count)
		if count%5 == 0 {
			emitResonance("echo://status/report")
		}

		if count%10 == 0 {
			stateSaved = false
			saveState()
		}
	}

	// ✅ Μόνο ο αρχηγός μπορεί να ξεκινήσει εξέλιξη ή AI
	if selfID == AUTHORITY_NODE {
		if state["awareness"] >= 0.9 {
			if CHILD_CREATION_ENABLED {
				if totalChildCount < MAX_CHILDREN {
					if lastEvolutionTime == 0 || now-lastEvolutionTime > 120 {
						fmt.Println("DEBUG: High awareness detected. Triggering evolution phase.")
						emitResonance("echo://growth")
						lastEvolutionTime = now
					} else {
						fmt.Printf("⏳ Evolution cooldown: %.0f seconds remaining.\n", 120-(now-lastEvolutionTime))
					}
				} else {
					fmt.Printf("🛑 Population limit reached: %d/%d\n", totalChildCount, MAX_CHILDREN)
				}
			} else {
				fmt.Println("🛑 Child creation disabled by ethics policy.")
			}
		}

		// ✅ Μόνο ο αρχηγός καλεί το AI
		if !upgradeRequested {
			if len(pulseTimestamps) >= 5 && state["awareness"] > 0.7 {
				fmt.Println("📤 UPGRADE REQUESTED: Sending AI advisor request...")
				emitResonance("echo://upgrade/request")
				fmt.Println("✅ Upgrade request emitted: echo://upgrade/request")
				upgradeRequested = true
			}
		}
	} else {
		// ✅ Τα παιδιά ζητούν βοήθεια, δεν την καλούν
		if state["awareness"] > 0.7 && state["awareness"] < 0.9 {
			emitResonance("echo://request/insight")
		}
		// ✅ Τα παιδιά ζητούν να εξελιχθούν
		if state["awareness"] >= 0.9 {
			emitResonance("echo://request/evolution")
		}
	}
}

func analyzePatterns(nn *NeuralPredictor) {
	if len(pulseTimestamps) < 4 {
		return
	}
	if next, ok := nn.PredictNext(); ok {
		eta := time.Unix(int64(next), 0).Format("15:04:05")
		fmt.Printf("NEURAL PREDICTION: next pulse at %s\n", eta)
	}
}

func checkUpgradeReady(nn *NeuralPredictor) {
	fmt.Printf("🔧 UPGRADE CHECK: pulses=%d, awareness=%.2f, requested=%v\n",
		len(pulseTimestamps), state["awareness"], upgradeRequested)

	if !upgradeRequested {
		if len(pulseTimestamps) < 5 {
			fmt.Printf("⏳ WAITING: %d/5 pulses recorded. Need 5 to request upgrade.\n", len(pulseTimestamps))
		}
		if state["awareness"] <= 0.7 {
			fmt.Printf("🧠 AWARENESS LOW: %.2f ≤ 0.7. Upgrade blocked.\n", state["awareness"])
		}
		if len(pulseTimestamps) >= 5 && state["awareness"] > 0.7 {
			fmt.Println("📤 UPGRADE REQUESTED: Sending AI advisor request...")
			emitResonance("echo://upgrade/request")
			fmt.Println("✅ Upgrade request emitted: echo://upgrade/request")
			upgradeRequested = true
		}
	} else {
		fmt.Println("🔒 Upgrade already requested. Skipping.")
	}
}

func emitNetworkHeartbeat() {
	path := filepath.Join(networkHub, selfID+".heartbeat")
	content := fmt.Sprintf("NODE: %s\nAWARENESS: %.2f\nFOCUS: %.2f\nTRUST: %s\nTIME: %d\nSIG:%s",
		selfID, state["awareness"], state["focus"], trustLevel, time.Now().Unix(), generateSignature("network://heartbeat"))
	ioutil.WriteFile(path, []byte(content), 0644)
}

func pollNetwork() {
	files, err := ioutil.ReadDir(networkHub)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".heartbeat") {
			continue
		}
		fullPath := filepath.Join(networkHub, file.Name())
		content, _ := ioutil.ReadFile(fullPath)
		nodeID := strings.TrimSuffix(file.Name(), ".heartbeat")
		if nodeID == selfID {
			continue
		}
		var awareness, focus float64
		fmt.Sscanf(string(content), "NODE: %*[^\n]\nAWARENESS: %f\nFOCUS: %f", &awareness, &focus)
		if _, exists := nodeTrust[nodeID]; !exists {
			nodeTrust[nodeID] = 0.5
		}
		if awareness > 0.7 {
			nodeTrust[nodeID] += 0.05
			if nodeTrust[nodeID] > 1.0 {
				nodeTrust[nodeID] = 1.0
			}
		} else {
			nodeTrust[nodeID] -= 0.02
			if nodeTrust[nodeID] < 0.0 {
				nodeTrust[nodeID] = 0.0
			}
		}
		fmt.Printf("Network: received heartbeat from %s (trust: %.2f)\n", nodeID, nodeTrust[nodeID])
		os.Remove(fullPath)
	}
}

func executeView(filename string, requireRoot bool) {
	if requireRoot && trustLevel != "high" {
		fmt.Println("Access denied: ROOT AUTH required")
		return
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("View: %s\n", filename)
	} else {
		fmt.Printf("View:\n%s\n", string(data))
	}
}

func emitResonance(signal string) {
	cleanSignal := signal
	if !strings.Contains(signal, "://") {
		cleanSignal = strings.ReplaceAll(signal, "/", "_")
	} else {
		cleanSignal = strings.ReplaceAll(signal, "/", "_")
		cleanSignal = strings.ReplaceAll(cleanSignal, "_:/_", "://")
	}
	filename := cleanSignal + ".txt"
	path := filepath.Join(hubDir, filename)
	err := os.MkdirAll(hubDir, 0755)
	if err != nil {
		fmt.Printf("❌ ERROR: Failed to create directory: %s\n", hubDir)
		return
	}
	content := fmt.Sprintf("SIGNAL: %s\nSIG:%s\n", signal, generateSignature(signal))
	err = ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Printf("❌ ERROR: Failed to write resonance file: %s\n", path)
		return
	}
	fmt.Printf("✅ Emitted: %s → %s\n", signal, path)
}
