import time
import os

PROXY_FILE = 'proxy.txt'
ROTATION_INTERVAL = 10  # seconds

def set_system_proxy(ip, port):
    os.system(f'netsh winhttp set proxy {ip}:{port}')
    print(f"[+] Proxy set to: {ip}:{port}")

def reset_proxy():
    os.system('netsh winhttp reset proxy')
    print("[!] Proxy reset to direct connection.")

def main():
    with open(PROXY_FILE, 'r') as f:
        proxies = [line.strip() for line in f if line.strip()]

    if not proxies:
        print("No proxies found in proxy.txt")
        return

    index = 0
    try:
        while True:
            proxy = proxies[index % len(proxies)]
            ip, port = proxy.split(':')
            set_system_proxy(ip, port)
            index += 1
            time.sleep(ROTATION_INTERVAL)
    except KeyboardInterrupt:
        reset_proxy()
        print("Exiting...")

if __name__ == '__main__':
    main()
