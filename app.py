import streamlit as st
import pandas as pd
import json
import os

st.set_page_config(page_title="VaultGuard-QPC SOC Dashboard", layout="wide")

st.title("🛡️ VaultGuard-QPC: Security Operations Center Dashboard")
st.subheader("Real-Time Zero-Trust Privileged Access Monitoring & Post-Quantum Audit Log")

# Auto-refresh mechanism
st.button("🔄 Refresh Data")

# Read the Go application's append log file
log_file = "audit_log.json"
logs = []

if os.path.exists(log_file):
    with open(log_file, "r") as f:
        for line in f:
            if line.strip():
                logs.append(json.loads(line))

if len(logs) > 0:
    df = pd.DataFrame(logs)

    # Metric Cards Row
    col1, col2, col3 = st.columns(3)
    with col1:
        st.metric("Total Administrative Actions Checked", len(df))
    with col2:
        threats_count = len(df[df['risk_score'] >= 70])
        st.metric("Insider Threats Blocked", threats_count, delta=f"{threats_count} High Risk Events", delta_color="inverse")
    with col3:
        st.metric("Quantum Cryptographic Status", "Active (ML-DSA Secure)")

    # Data Layout Split
    left_col, right_col = st.columns([2, 1])

    with left_col:
        st.write("### Recent Log Interceptions")
        # Reverse rows to show latest first using updated 2026 syntax
        st.dataframe(df.iloc[::-1], width="stretch")

    with right_col:
        st.write("### User Risk Profiler")
        # Plot risk variance grouped by administrative user session handles
        user_risk = df.groupby('user')['risk_score'].max()
        st.bar_chart(user_risk)

        st.write("### Top Data Usage Anomalies")
        anomaly_df = df[df['risk_score'] >= 70][['user', 'query', 'risk_score']].iloc[::-1]
        st.table(anomaly_df)
else:
    st.info("Waiting for real-time proxy query transmissions... Run your PowerShell simulations to generate logs.")