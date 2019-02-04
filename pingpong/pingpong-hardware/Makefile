# Created by: girgen
# $FreeBSD$

PORTNAME=	pingpong-hardware
PORTVERSION=	0.6
CATEGORIES=	pingpong sysutils
MASTER_SITES=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	Meta-port for PING PONG hardware monitoring tools

LICENSE=	BSD2CLAUSE

RUN_DEPENDS= \
	${LOCALBASE}/libexec/nagios/check_smartmon:net-mgmt/nagios-check_smartmon \
	zfSnap:sysutils/zfsnap \
	amrstat:sysutils/amrstat \
	cciss_vol_status:sysutils/cciss_vol_status \
	${LOCALBASE}/etc/rc.d/microcode_update:sysutils/devcpu-data \
	p5-HTML-Table>0:www/p5-HTML-Table \
	arcconf:sysutils/arcconf

USES=		metaport

NO_BUILD=	yes
NO_MTREE=	yes

.include <bsd.port.mk>
