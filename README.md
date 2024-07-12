# Study golang

# Visual Studio Code 설치
## Windows
1. [VS Code 다운로드 페이지](https://code.visualstudio.com/)로 이동합니다.
2. Windows용 설치 프로그램을 다운로드합니다.
3. 다운로드한 설치 프로그램을 실행하고 지시에 따라 설치합니다.

## macOS
1. [VS Code 다운로드 페이지](https://code.visualstudio.com/)로 이동합니다.
2. macOS용 .dmg 파일을 다운로드합니다.
3. 다운로드한 .dmg 파일을 열고, VS Code 아이콘을 Applications 폴더로 드래그하여 설치합니다.

## Linux

1. [VS Code 다운로드 페이지](https://code.visualstudio.com/)로 이동합니다.
2. 배포판에 맞는 설치 파일을 다운로드합니다 (Debian/Ubuntu용 .deb 파일 또는 Red Hat/Fedora/SUSE용 .rpm 파일).
3. 터미널을 열고, 다운로드한 파일이 있는 디렉토리로 이동한 다음, 다음 명령을 실행하여 설치합니다:

> Debian/Ubuntu:

    sudo dpkg -i code_*.deb
    sudo apt-get install -f  # dependencies가 필요할 경우

> Red Hat/Fedora/SUSE:

    sudo rpm -i code-*.rpm

# Go 확장 기능 설치

1. VS Code를 실행합니다.
2. 사이드바에서 확장 아이콘을 클릭하거나, Ctrl+Shift+X (Windows/Linux) 또는 Cmd+Shift+X (macOS)를 눌러 확장 창을 엽니다.
3. 검색 창에 Go를 입력하고, Go (by Go Team at Google) 확장을 찾아 설치합니다.

# Go 개발 환경 설정
## Go 언어 설치

1. [Go 언어 다운로드 페이지](https://go.dev/dl/)로 이동하여, 운영 체제에 맞는 설치 파일을 다운로드합니다.

2. 설치 프로그램을 실행하고 지시에 따라 Go 언어를 설치합니다.

## Go 환경 변수 설정 (Windows)

1. 설치 완료 후, Windows + R을 눌러 실행 창을 열고 cmd를 입력하여 명령 프롬프트를 엽니다.
2. 명령 프롬프트에 go version을 입력하여 Go가 정상적으로 설치되었는지 확인합니다.
3. Go가 정상적으로 설치되었으면, 환경 변수를 설정합니다.

> - 시스템 속성 -> 고급 -> 환경 변수로 이동합니다.
> - 시스템 변수 섹션에서 Path를 선택하고 편집을 클릭합니다.
> - 새 항목으로 Go의 bin 디렉토리 경로를 추가합니다 (예: C:\Go\bin).
> - 새 시스템 변수 GOPATH를 만들고, Go 작업 디렉토리 경로를 설정합니다 (예: C:\Users\<YourUsername>\go).

## Go 환경 변수 설정 (macOS/Linux)

- 터미널을 엽니다.
- 홈 디렉토리에 있는 쉘 프로파일 파일을 엽니다 (.bashrc, .bash_profile, .zshrc 등).
- 다음 내용을 추가합니다

    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
- 파일을 저장하고 쉘을 다시 로드합니다:

    source ~/.bashrc # or appropriate file

# VS Code에서 Go 프로젝트 시작

- VS Code에서 Ctrl+Shift+P (Windows/Linux) 또는 Cmd+Shift+P (macOS)를 눌러 명령 팔레트를 엽니다.
- Go: Install/Update Tools를 입력하고 선택하여 필요한 Go 도구들을 설치합니다.
- 새로운 Go 프로젝트를 시작하려면, 파일 -> 열기 -> 폴더를 선택하고 프로젝트 디렉토리를 선택합니다.
- main.go 파일을 만들고 Go 코드를 작성합니다. 예:

    package main
    import "fmt"
    func main() {
        fmt.Println("Hello, World!")
    }

- 터미널을 열고(`Ctrl+``) 다음 명령을 실행하여 프로그램을 빌드하고 실행합니다:

    go run main.go

이제 VS Code와 Go 개발 환경이 설정되었습니다. 이 설정으로 Go 개발을 시작할 수 있습니다. VS Code는 코드 자동 완성, 디버깅, linting, 테스트 등의 강력한 기능을 제공하여 Go 개발을 더 쉽고

효율적으로 만들어줍니다.
