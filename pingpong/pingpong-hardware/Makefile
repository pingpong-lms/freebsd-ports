# Created by: girgen
#
PORTNAME=	pingpong-hardware
PORTVERSION=	0.5
CATEGORIES=	pingpong sysutils
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for PING PONG hardware monitoring tools

NO_BUILD=	yes
NO_MTREE=	yes

RUN_DEPENDS= \
	${LOCALBASE}/libexec/nagios/check_smartmon:net-mgmt/nagios-check_smartmon \
	zfSnap:sysutils/zfsnap \
	amrstat:sysutils/amrstat \
	cciss_vol_status:sysutils/cciss_vol_status \
	${LOCALBASE}/etc/rc.d/microcode_update:sysutils/devcpu-data \
	arcconf:sysutils/arcconf

do-install: build
	${CHMOD} u+s ${LOCALBASE}/sbin/smartctl ${LOCALBASE}/bin/cciss_vol_status

.include <bsd.port.mk>
