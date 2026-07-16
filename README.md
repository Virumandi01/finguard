# finguard
financial guard
 to run the file :
 go run main.go


 to est run these in the another new terminal :
            shows the test case login with a=other ip and try to retrivee many data set apart for their quota average and other sinerao :


Scenario A: Safe, Baseline Behavior
Simulate admin_alex running a low-risk, filtered database query inside working hours from his approved workstation IP (192.168.1.25).

PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()
$data = [System.Text.Encoding]::UTF8.GetBytes("admin_alex|192.168.1.25|SELECT id FROM balances WHERE id = 450")
$stream.Write($data, 0, $data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()
$client.Close()
Expected Output in Terminal 2: SUCCESS: Transaction query executed smoothly.

Expected Output in Terminal 1 (The Go Window): You will see an immutable log entry with a Risk Score: 0 or low value, signed with a post-quantum cryptographic signature.



Scenario B: Unregistered User (Insider Threat Detect)
Simulate a completely unknown user attempting to execute commands on the database.
PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()
$data = [System.Text.Encoding]::UTF8.GetBytes("unknown_attacker|192.168.1.25|SELECT * FROM users")
$stream.Write($data, 0, $data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()
$client.Close()
Expected Output in Terminal 2: ACCESS DENIED: Session terminated. Violation: CRITICAL: Unregistered administrative credentials detected

Expected Output in Terminal 1: An immediate alert printout with a Risk Score: 95 along with a cryptographic signature proof.


Scenario C: Data Exfiltration (SQL Data Monitoring & IP Jump)
Simulate admin_alex logging in from a completely different network IP address (203.0.113.5) outside the bank's network profile and attempting a bulk dump of the entire financial accounts table.

PowerShell
$client = New-Object System.Net.Sockets.TcpClient("localhost", 8080)
$stream = $client.GetStream()
$data = [System.Text.Encoding]::UTF8.GetBytes("admin_alex|203.0.113.5|SELECT * FROM core_bank_accounts")
$stream.Write($data, 0, $data.Length)
$reader = New-Object System.IO.StreamReader($stream)
$reader.ReadLine()
$client.Close()
Expected Output in Terminal 2: ACCESS DENIED: Session terminated. Violation: IP network mesh mismatch, Bulk table dump sequence scanned, Targeting high-risk financial registries

Expected Output in Terminal 1: A drop log showing the calculation shot straight past your threshold to Risk Score: 95 (or 100 depending on the current time window check).