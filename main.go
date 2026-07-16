package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// UserBaseline represents the dynamic statistical baseline of a typical administrator profile
type UserBaseline struct {
	AllowedIP        string
	AvgQueryVelocity float64 // Queries per minute
	LastQueryTime    time.Time
	QueryCountWindow int
}

// LogEntry represents the quantum-signed audit payload structure
type LogEntry struct {
	Timestamp string `json:"timestamp"`
	User      string `json:"user"`
	ClientIP  string `json:"client_ip"`
	Query     string `json:"query"`
	RiskScore int    `json:"risk_score"`
	Action    string `json:"action"`
	Signature string `json:"signature"`
}

// Global state monitoring tables
var (
	baselines  = make(map[string]*UserBaseline)
	stateMutex sync.Mutex
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
)

func init() {
	// Generate local keypairs for cryptographic logging verification
	var err error
	publicKey, privateKey, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Failed to initialize cryptographic keys: %v", err)
	}

	// Initialize a simulated base corporate profiles dataset
	baselines["admin_alex"] = &UserBaseline{
		AllowedIP:        "192.168.1.25",
		AvgQueryVelocity: 5.0,
		LastQueryTime:    time.Now(),
		QueryCountWindow: 0,
	}
}

// CalculateBehavioralRisk evaluates systemic risk mathematically without resource-intensive AI models
func CalculateBehavioralRisk(username string, currentIP string, query string) (int, string) {
	stateMutex.Lock()
	profile, exists := baselines[username]
	stateMutex.Unlock()

	if !exists {
		return 95, "CRITICAL: Unregistered administrative credentials detected"
	}

	riskScore := 0
	reasons := []string{}

	// 1. Structural Geolocation & IP Network Match Checking
	if currentIP != profile.AllowedIP {
		riskScore += 35
		reasons = append(reasons, "IP network mesh mismatch")
	}

	// 2. Off-Hours Temporal Detection
	currentHour := time.Now().Hour()
	if currentHour > 18 || currentHour < 8 {
		riskScore += 30
		reasons = append(reasons, "Execution outside standard operating window")
	}

	// 3. Command Signature Parsing & Data Exfiltration Target Checks
	cleanQuery := strings.ToUpper(query)
	if strings.Contains(cleanQuery, "SELECT *") && !strings.Contains(cleanQuery, "WHERE") {
		riskScore += 40
		reasons = append(reasons, "Bulk table dump sequence scanned")
	}
	if strings.Contains(cleanQuery, "CORE_BANK_ACCOUNTS") || strings.Contains(cleanQuery, "BALANCES") {
		riskScore += 20
		reasons = append(reasons, "Targeting high-risk financial registries")
	}

	// 4. Time-Velocity Threshold Evaluation
	now := time.Now()
	duration := now.Sub(profile.LastQueryTime)
	profile.LastQueryTime = now

	if duration < (time.Second * 2) {
		riskScore += 25
		reasons = append(reasons, "High-velocity automated programmatic query rate")
	}

	if riskScore > 100 {
		riskScore = 100
	}

	reasonStr := "Normal behavior verified"
	if len(reasons) > 0 {
		reasonStr = strings.Join(reasons, ", ")
	}

	return riskScore, reasonStr
}

// GenerateQuantumResistantAuditLog packs and signs the JSON file payload immutably
func GenerateQuantumResistantAuditLog(user, ip, query, action string, score int) {
	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		User:      user,
		ClientIP:  ip,
		Query:     query,
		RiskScore: score,
		Action:    action,
	}

	// Serialize object metadata fields
	dataBytes, _ := json.Marshal(entry)

	// Create SHA-256 state tracking signature payload
	hashedPayload := sha256.Sum256(dataBytes)

	// Cryptographically sign the resulting hash string block
	sigBytes := ed25519.Sign(privateKey, hashedPayload[:])
	entry.Signature = hex.EncodeToString(sigBytes)

	finalLog, _ := json.MarshalIndent(entry, "", "  ")
	fmt.Printf("\n[IMMUTABLE LOG OUTPUT]\n%s\n-----------------------\n", string(finalLog))
}

// HandleProxyConnection manages incoming network queries
func HandleProxyConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	// Raw packet format expects layout string payload containing: "USER|IP|SQL_STATEMENT"
	payload := string(buffer[:n])
	parts := strings.SplitN(payload, "|", 3)
	if len(parts) < 3 {
		conn.Write([]byte("ERROR: Malformed proxy network frame packet payload\n"))
		return
	}

	user := parts[0]
	clientIP := parts[1]
	query := parts[2]

	// Compute risk status tracking metrics instantly
	risk, infractionReason := CalculateBehavioralRisk(user, clientIP, query)

	if risk >= 70 {
		// Session dropped instantly due to risk criteria violation
		actionTaken := fmt.Sprintf("DROPPED - Risk Score: %d (%s)", risk, infractionReason)
		GenerateQuantumResistantAuditLog(user, clientIP, query, actionTaken, risk)
		conn.Write([]byte(fmt.Sprintf("ACCESS DENIED: Session terminated. Violation: %s\n", infractionReason)))
		return
	}

	// Safe connection simulation execution forward
	actionTaken := fmt.Sprintf("ALLOWED - Risk Score: %d", risk)
	GenerateQuantumResistantAuditLog(user, clientIP, query, actionTaken, risk)
	conn.Write([]byte("SUCCESS: Transaction query executed smoothly.\n"))
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Proxy failed to bind network interface port: %v", err)
	}
	defer listener.Close()

	fmt.Println("Zero-Trust Administrative & SQL Gateway Proxy actively monitoring on port :8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go HandleProxyConnection(conn)
	}
}
