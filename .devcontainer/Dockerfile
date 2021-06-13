# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.166.1/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.16, 1.15
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# Or your actual UID, GID on Linux if not the default 1000
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=${USER_UID}
ARG TZ=America/Los_Angeles

# The DOCKER_COMPOSE_SHA can be found at https://github.com/docker/compose/releases, inside the corresponding .sha256
# file of the current release.
ARG DOCKER_COMPOSE_SHA=f3f10cf3dbb8107e9ba2ea5f23c1d2159ff7321d16f0a23051d68d8e2547b323
ARG DOCKER_COMPOSE_VERSION=1.29.2

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su ${USERNAME} -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

RUN set -e && \
# Install the official Docker repo and add the key for it
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    echo "deb https://download.docker.com/linux/debian buster stable" | tee -a /etc/apt/sources.list.d/docker.list && \
    apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
        software-properties-common \
        build-essential \
        git \
        htop \
        openssh-server \
        gnupg-agent \
        zsh \
        iputils-ping \
        traceroute \
        dnsutils \
        net-tools \
        host \
        unzip \
        jq \
        vim-tiny \
        libyaml-0-2 \
        zlib1g-dev \
        libcurl4-openssl-dev \
        sudo \
        less \
        file \
        docker-ce-cli \
        sqlite3 && \
# local timezone
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
## docker-compose
    curl -L -O https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-Linux-x86_64 && \
    echo "${DOCKER_COMPOSE_SHA} docker-compose-Linux-x86_64" | sha256sum -c - && \
    mv docker-compose-Linux-x86_64 /usr/bin/docker-compose && \
    chmod +sx /usr/bin/docker-compose && \
    mv /usr/bin/docker-compose /usr/bin/docker-compose.orig && \
# clean up apt
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN set -e && \
# Create a non-root user to use if preferred - see https://aka.ms/vscode-remote/containers/non-root-user.
    userdel vscode && \
    groupadd --gid $USER_GID $USERNAME && \
    useradd -s /usr/bin/zsh --uid $USER_UID --gid $USER_GID -m $USERNAME && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME && \
    # store vscode-server files in the contaienr root 
    mkdir /.vscode-server && \
    chown yarmand.yarmand /.vscode-server && \
    mkdir /.vscode-server-insiders && \
    chown yarmand.yarmand /.vscode-server-insiders


# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su ${USERNAME} -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1