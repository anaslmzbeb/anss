#!/bin/bash

# Flush all existing rules
iptables -F
iptables -P INPUT DROP
iptables -P FORWARD DROP
iptables -P OUTPUT ACCEPT

# Allow localhost and existing connections
iptables -A INPUT -i lo -j ACCEPT
iptables -A INPUT -m conntrack --ctstate ESTABLISHED,RELATED -j ACCEPT

# Allow Morocco IPs
for ip in $(cat ma.zone); do
  iptables -A INPUT -s $ip -j ACCEPT
done

# Allow Algeria IPs
for ip in $(cat dz.zone); do
  iptables -A INPUT -s $ip -j ACCEPT
done

# Allow Tunisia IPs
for ip in $(cat tn.zone); do
  iptables -A INPUT -s $ip -j ACCEPT
done

# Allow Italy IPs
for ip in $(cat it.zone); do
  iptables -A INPUT -s $ip -j ACCEPT
done
