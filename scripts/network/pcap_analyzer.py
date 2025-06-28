#!/usr/bin/env python3

import sys
import json
import os
from scapy.all import rdpcap, IP, TCP, UDP, ICMP, ARP

def analyze_pcap(pcap_file):
    try:
        # Read the PCAP file
        packets = rdpcap(pcap_file)
        
        # Initialize analysis results
        analysis = {
            "total_packets": len(packets),
            "protocols": {},
            "ip_addresses": {
                "source": {},
                "destination": {}
            },
            "ports": {
                "tcp": {},
                "udp": {}
            },
            "top_conversations": [],
            "suspicious_activity": []
        }
        
        conversations = {}
        
        for packet in packets:
            # Protocol analysis
            if packet.haslayer(IP):
                analysis["protocols"]["IP"] = analysis["protocols"].get("IP", 0) + 1
                
                src_ip = packet[IP].src
                dst_ip = packet[IP].dst
                
                # Track IP addresses
                analysis["ip_addresses"]["source"][src_ip] = analysis["ip_addresses"]["source"].get(src_ip, 0) + 1
                analysis["ip_addresses"]["destination"][dst_ip] = analysis["ip_addresses"]["destination"].get(dst_ip, 0) + 1
                
                # Track conversations
                conv_key = f"{src_ip}:{dst_ip}"
                conversations[conv_key] = conversations.get(conv_key, 0) + 1
                
                # TCP analysis
                if packet.haslayer(TCP):
                    analysis["protocols"]["TCP"] = analysis["protocols"].get("TCP", 0) + 1
                    src_port = packet[TCP].sport
                    dst_port = packet[TCP].dport
                    
                    analysis["ports"]["tcp"][dst_port] = analysis["ports"]["tcp"].get(dst_port, 0) + 1
                    
                    # Check for suspicious ports
                    suspicious_ports = [22, 23, 80, 443, 21, 25, 53, 135, 139, 445]
                    if dst_port in suspicious_ports:
                        analysis["suspicious_activity"].append({
                            "type": "suspicious_port",
                            "src": src_ip,
                            "dst": dst_ip,
                            "port": dst_port,
                            "protocol": "TCP"
                        })
                
                # UDP analysis
                elif packet.haslayer(UDP):
                    analysis["protocols"]["UDP"] = analysis["protocols"].get("UDP", 0) + 1
                    dst_port = packet[UDP].dport
                    analysis["ports"]["udp"][dst_port] = analysis["ports"]["udp"].get(dst_port, 0) + 1
                
                # ICMP analysis
                elif packet.haslayer(ICMP):
                    analysis["protocols"]["ICMP"] = analysis["protocols"].get("ICMP", 0) + 1
            
            # ARP analysis
            elif packet.haslayer(ARP):
                analysis["protocols"]["ARP"] = analysis["protocols"].get("ARP", 0) + 1
        
        # Top conversations
        analysis["top_conversations"] = [
            {"conversation": k, "packets": v} 
            for k, v in sorted(conversations.items(), key=lambda x: x[1], reverse=True)[:10]
        ]
        
        return {
            "success": True,
            "analysis": analysis,
            "tool": "pcap_analyzer"
        }
        
    except Exception as e:
        return {
            "success": False,
            "error": f"PCAP analysis failed: {str(e)}"
        }

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print(json.dumps({
            "success": False,
            "error": "Usage: python pcap_analyzer.py <pcap_file>"
        }))
        sys.exit(1)
    
    pcap_file = sys.argv[1]
    
    if not os.path.exists(pcap_file):
        print(json.dumps({
            "success": False,
            "error": "PCAP file not found"
        }))
        sys.exit(1)
    
    result = analyze_pcap(pcap_file)
    print(json.dumps(result, indent=2))