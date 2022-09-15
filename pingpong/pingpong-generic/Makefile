# Created by: girgen
#
PORTNAME=	pingpong-generic
PORTVERSION=	0.23
CATEGORIES=	pingpong sysutils
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for generic ppab server deps

NO_BUILD=	yes

USES=		perl5 gmake pgsql python autoreconf
WANT_PGSQL=	client
USE_OPENLDAP=	yes

PLIST_FILES=	share/vim/vimrc \
		etc/periodic/daily/pp_pkg_audit etc/newsyslog.conf.d/pp_pkg_audit.conf
PORTDATA=	installed.stamp

RUN_DEPENDS= \
	${BUILD_DEPENDS} \
	bash:shells/bash \
	${LOCALBASE}/share/bash-completion/bash_completion:shells/bash-completion \
	zsh:shells/zsh \
	vim:editors/vim \
	tree:sysutils/tree \
	curl:ftp/curl \
	mosh:net/mosh \
	lftp:ftp/lftp \
	rsync:net/rsync \
	unzip:archivers/unzip \
	svn:devel/subversion \
	python:lang/python \
	python2.7:lang/python27 \
	git:devel/git \
	p5-DBD-Pg>0:databases/p5-DBD-Pg \
	p5-Switch>0:lang/p5-Switch \
	p5-Proc-ProcessTable>0:devel/p5-Proc-ProcessTable \
	p5-Crypt-SSLeay>0:security/p5-Crypt-SSLeay \
	p5-Net-DNS>0:dns/p5-Net-DNS \
	nrpe3:net-mgmt/nrpe3 \
	collectd:net-mgmt/collectd5 \
	telegraf:net-mgmt/telegraf \
	ossec-hids-agent-config>=3.7.0:security/ossec-hids-agent-config \
	snmpd:net-mgmt/net-snmp \
	lsof:sysutils/lsof \
	dig:dns/bind-tools \
	lynx:www/lynx \
	gpg2:security/gnupg \
	doas:security/doas \
	sudo:security/sudo \
	nslcd:net/nss-pam-ldapd-sasl \
	${LOCALBASE}/lib/pam_mkhomedir.so:security/pam_mkhomedir \
	xsltproc:textproc/libxslt \
	xauth:x11/xauth \
	ag:textproc/the_silver_searcher \
	rg:textproc/ripgrep \
	fd:sysutils/fd \
	tmux:sysutils/tmux \
	screen:sysutils/screen

.include <bsd.port.options.mk>
.if ${OSVERSION} >= 1200000
RUN_DEPENDS+=	rcs:devel/rcs
.endif

configure:
	${DO_NADA}

do-install:
	${MKDIR} ${STAGEDIR}${PREFIX}/share/vim && \
	${INSTALL_DATA} ${FILESDIR}/vimrc ${STAGEDIR}${PREFIX}/share/vim
	${MKDIR} ${STAGEDIR}${PREFIX}/etc/periodic/daily ${STAGEDIR}${PREFIX}/etc/newsyslog.conf.d
	${INSTALL_SCRIPT} ${FILESDIR}/pp_pkg_audit ${STAGEDIR}${PREFIX}/etc/periodic/daily/pp_pkg_audit
	${INSTALL_DATA} ${FILESDIR}/pp_pkg_audit.conf ${STAGEDIR}${PREFIX}/etc/newsyslog.conf.d/pp_pkg_audit.conf
	${MKDIR} ${STAGEDIR}${DATADIR} && \
	${TOUCH} ${STAGEDIR}${DATADIR}/${PORTDATA}

.include <bsd.port.mk>
