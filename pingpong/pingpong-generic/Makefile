# Created by: girgen
#
PORTNAME=	pingpong-generic
PORTVERSION=	0.16
CATEGORIES=	pingpong sysutils
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for generic ppab server deps

NO_BUILD=	yes

USES=		perl5 gmake pgsql python
WANT_PGSQL=	client
USE_AUTOTOOLS=	autoconf aclocal automake
USE_OPENLDAP=	yes
WANT_OPENLDAP_SASL=yes

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
	python:lang/python \
	curl:ftp/curl \
	mosh:net/mosh \
	lftp:ftp/lftp \
	rsync:net/rsync \
	unzip:archivers/unzip \
	svn:devel/subversion \
	git:devel/git \
	p5-DBD-Pg>0:databases/p5-DBD-Pg \
	p5-Switch>0:lang/p5-Switch \
	p5-Proc-ProcessTable>0:devel/p5-Proc-ProcessTable \
	p5-Crypt-SSLeay>0:security/p5-Crypt-SSLeay \
	p5-Net-DNS>0:dns/p5-Net-DNS \
	nrpe3:net-mgmt/nrpe3 \
	collectd:net-mgmt/collectd5 \
	telegraf:net-mgmt/telegraf \
	snmpd:net-mgmt/net-snmp \
	lsof:sysutils/lsof \
	dig:dns/bind-tools \
	lynx:www/lynx \
	gpg2:security/gnupg \
	sudo:security/sudo \
	nslcd:net/nss-pam-ldapd-sasl \
	${LOCALBASE}/lib/pam_mkhomedir.so:security/pam_mkhomedir \
	xsltproc:textproc/libxslt \
	xauth:x11/xauth \
	ag:textproc/the_silver_searcher \
	tmux:sysutils/tmux \
	screen:sysutils/screen

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
