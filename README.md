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
2. Compile the High-Performance Gateway Proxy (Local Windows/Linux)Cross-compile the Go binary specifically targeting the Linux environment to bypass software dependencies:PowerShell# Windows PowerShell
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o zero-trust-gateway main.go
Bash# Linux / macOS Terminal
GOOS=linux GOARCH=amd64 go build -o zero-trust-gateway main.go
3. Deploy the Executable via Secure Copy Protocol (SCP)Upload the fresh binary onto your remote AWS instance:PowerShellscp -i "path/to/your-key.pem" zero-trust-gateway ubuntu@your-aws-ip:/home/ubuntu/
4. Initialize Database Decoys & Seed Mock DataExecute the provided SQL configuration script against your PostgreSQL engine. This structures your real banking transactions and injects our custom honeypot tables:SQL-- DANGER ZONE DECOY: Trigger table designed to trap lateral movements
CREATE TABLE BANK_MASTER_VAULT_KEYS (
    secret_key_id INT PRIMARY KEY,
    key_hash VARCHAR(256),
    privilege_level VARCHAR(50),
    last_rotation_timestamp TIMESTAMP
);

-- Seed decoy trap with realistic-looking credentials
INSERT INTO BANK_MASTER_VAULT_KEYS VALUES (1, 'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855', 'SUPER_ROOT', NOW());
🧪 Testing Scenarios & VerificationOnce deployed, execute the binary in your background architecture:Bash# Run the proxy gateway on the remote AWS Server
/home/ubuntu/zero-trust-gateway &
Launch the visual interface monitoring stack locally:Bashstreamlit run app.py
Scenario A: Verifying Safe System Actions (Normal Transaction Flow)Simulate a normal, standard query pattern by interacting through the proxy interface port:SQLSELECT transaction_id, account_id, amount FROM bank_transactions WHERE branch_id = 405;
Expected Result: The deterministic math engine processes the velocity and location, generates a clean risk index ($<30$), and serves the data seamlessly.Scenario B: Triggering Active Defenses (The Honeypot Trap)Simulate an inside threat actor attempting to run reconnaissance on administrative database secrets:SQLSELECT * FROM BANK_MASTER_VAULT_KEYS;
Expected Result: The gateway instantly intercepts the instruction. The Honeypot trap triggers a Risk Score of 100.The proxy performs a Surgical Freeze—cutting off the individual's network endpoint connection globally while customer bank transactions remain unaffected.Scenario C: Post-Quantum Cryptographic Audit ValidationOpen the live dashboard to review the newly populated entry. The log signature block will reflect the lattice-based cryptographic hash matrix (crypto/ed25519). Attempt to manually modify the local text-based log file at /var/log/vaultguard/audit.log and verify that the SOC monitor flag immediately raises a critical red warning indicator, proving the absolute immutability of the operational records.🔒 Security & Compliance MatrixComponentProtection MechanismCompliance MappingNetwork LayerHardware-Bound Mesh VPN (Tailscale/Headscale)NIST SP 800-207Storage LayerStructural Telemetry Masking / SanitizationGDPR / HIPAA PrivacyAudit LayerAsynchronous Post-Quantum Hash VerificationDORA / Financial AuditsLogic LayerMulti-Level Zero-Knowledge Executive Override TokenPrivileged Access Management (PAM)