function collectd() {
    local src="$DIR/addons/collectd/$COLLECTD_VERSION"
    
    if ! systemctl list-units | grep -q collectd; then
        printf "${YELLOW}Installing collectd\n"
        case "$LSB_DIST" in
        ubuntu)
            export DEBIAN_FRONTEND=noninteractive
            dpkg --install --no-install-recommends --force-depends-version ${src}/ubuntu-${DIST_VERSION}/archives/*.deb
            ;;
        centos|rhel|amzn)
            rpm --upgrade --force --nodeps ${src}/rhel-${DIST_VERSION}/archives/*.rpm
            ;;
        *)
            printf"${YELLOW}Unsupported OS for collectd installation\n"
            ;;
        esac

        collectd_service
    else
        printf "${YELLOW} Collectd is running, skipping installation\n"
    fi
}

function collectd_service() {
    if ! systemctl -q is-active collectd; then
        systemctl start collectd
    fi

    if ! systemctl -q is-enabled collectd; then
        systemctl enable collectd
    fi
}

function collectd_join() {
    collectd
}
