# VaultGuard-QPC 🛡️⚛️
> **Active Insider Threat Neutralization & Post-Quantum Architectural Defense for Enterprise FinTech**

VaultGuard-QPC is an enterprise-grade, high-performance zero-trust interceptor proxy designed to eliminate malicious insider activities, credential leaks, and data exfiltration at the database layer. By moving away from passive log collection, VaultGuard-QPC introduces **Active Deception (Honeypot Decoys)**, **Surgical Blast-Radius Containment**, and **Post-Quantum Cryptographic (QPC)** audit trails.

---

## 🏗️ Core Architecture Overview

VaultGuard-QPC splits its enforcement and operational layers into isolated modules to maintain low-latency database execution and ensure high availability.

1. **The Interceptor Gateway (`zero-trust-gateway`):** A lightweight, compiled Go binary that acts as a reverse proxy directly in front of your database instances. It houses the deterministic risk scoring matrices and honeypot traps.
2. **Zero-Knowledge Override Module:** An isolated administrative daemon accessible only by top-tier executives via hardware keys to issue one-time bypass tokens or defreeze locked states.
3. **On-Demand AI Behavioral Analyzer:** A completely optional, standalone Python microservice that reads *sanitized* telemetry logs and utilizes the Gemini API for forensic anomaly summaries without exposing raw banking data.
4. **The SOC Dashboard:** A real-time monitoring center built in Python (Streamlit & Pandas) parsing the immutable QPC audit logs.

---

## 📋 Prerequisites & Ecosystem Requirements

Before proceeding with deployment, ensure your target environments meet the following configurations:

* **Production Target / Staging:** AWS EC2 instance running Ubuntu Server 22.04 LTS or higher.
* **Network Infrastructure:** Active **Tailscale / Headscale** mesh network configured on both the database host machine and the administrator terminal node.
* **Database Engine:** PostgreSQL (or SQLite/MySQL) hosting your production tables.
* **Local Machine Development Tools:** Go v1.21+ (for cross-compilation), Python 3.11+, and an active Google Gemini API key (optional).

---

## 🚀 Deployment Step-by-Step

### 1. Set Up the Target Environment (AWS Ubuntu)
SSH into your AWS Ubuntu Server instance inside your private network mesh and pull up the target directories:
```bash
sudo mkdir -p /etc/vaultguard /var/log/vaultguard
sudo chown -R ubuntu:ubuntu /etc/vaultguard /var/log/vaultguard
2. Compile the High-Performance Gateway Proxy (Local Windows/Linux)
Cross-compile the Go binary specifically targeting the Linux environment to bypass software dependencies:

PowerShell
# Windows PowerShell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o zero-trust-gateway main.go
Bash
# Linux / macOS Terminal
GOOS=linux GOARCH=amd64 go build -o zero-trust-gateway main.go
3. Deploy the Executable via Secure Copy Protocol (SCP)
Upload the fresh binary onto your remote AWS instance:

PowerShell
scp -i "path/to/your-key.pem" zero-trust-gateway ubuntu@your-aws-ip:/home/ubuntu/
4. Initialize Database Decoys & Seed Mock Data
Execute the provided SQL configuration script against your PostgreSQL engine. This structures your real banking transactions and injects our custom honeypot tables:

SQL
-- DANGER ZONE DECOY: Trigger table designed to trap lateral movements
CREATE TABLE BANK_MASTER_VAULT_KEYS (
    secret_key_id INT PRIMARY KEY,
    key_hash VARCHAR(256),
    privilege_level VARCHAR(50),
    last_rotation_timestamp TIMESTAMP
);

-- Seed decoy trap with realistic-looking credentials
INSERT INTO BANK_MASTER_VAULT_KEYS VALUES (1, 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855', 'SUPER_ROOT', NOW());
🧪 Testing Scenarios & Sample Commands
To execute tests, launch the system modules across your environment windows.

Start the compiled proxy daemon on your remote AWS server terminal (Terminal 1):

Bash
/home/ubuntu/zero-trust-gateway
Launch the visual SOC tracking analytics panel on your local administrative station terminal:

Bash
streamlit run app.py
Open a local terminal (Terminal 2) to simulate administrative pipeline queries using the standard PowerShell socket pipelines outlined below.

🟢 Scenario A: Safe, Baseline Behavior
Simulate admin_alex running a low-risk, filtered database query inside working hours from his approved workstation IP (192.168.1.25).

Run this command:

PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()$data = [System.Text.Encoding]::UTF8.GetBytes("admin_alex|192.168.1.25|SELECT id FROM balances WHERE id = 450")
$stream.Write($data, 0,$data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()$client.Close()
Expected Output in Terminal 2: SUCCESS: Transaction query executed smoothly.

Expected Output in Terminal 1 (The Go Window): You will see an immutable log entry print out with a Risk Score: 0 or low value, signed with a post-quantum cryptographic signature.

🟡 Scenario B: Unregistered User (Insider Threat Detect)
Simulate a completely unknown user attempting to execute commands on the database.

Run this command:

PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()$data = [System.Text.Encoding]::UTF8.GetBytes("unknown_attacker|192.168.1.25|SELECT * FROM users")
$stream.Write($data, 0,$data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()$client.Close()
Expected Output in Terminal 2: ACCESS DENIED: Session terminated. Violation: CRITICAL: Unregistered administrative credentials detected

Expected Output in Terminal 1: An immediate alert printout with a Risk Score: 95 along with a cryptographic signature proof.

🔴 Scenario C: Data Exfiltration (SQL Data Monitoring & IP Jump)
Simulate admin_alex logging in from a completely different network IP address (203.0.113.5) outside the bank's network profile and attempting a bulk dump of the entire financial accounts table.

Run this command:

PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()$data = [System.Text.Encoding]::UTF8.GetBytes("admin_alex|203.0.113.5|SELECT * FROM core_bank_accounts")
$stream.Write($data, 0,$data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()$client.Close()
Expected Output in Terminal 2: ACCESS DENIED: Session terminated. Violation: IP network mesh mismatch, Bulk table dump sequence scanned, Targeting high-risk financial registries

Expected Output in Terminal 1: The proxy drops the pipeline instantly, writes a Risk Score: 100 signature record, blocks the account from further sessions, and routes an isolated notice payload to the Streamlit dashboard screen.