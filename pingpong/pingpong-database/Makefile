# Created by: girgen
#
PORTNAME=	pingpong-database
PORTVERSION=	0.6
CATEGORIES=	pingpong databases
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for PING PONG application packages

NO_BUILD=	yes
NO_MTREE=	yes

USES=		pgsql
WANT_PGSQL_VER=	96
WANT_PGSQL=	server contrib

RUN_DEPENDS=	${LOCALBASE}/share/pingpong-generic/installed.stamp:pingpong/pingpong-generic \
		${LOCALBASE}/lib/postgresql/powa.so:databases/powa-archivist \
		pgbadger:databases/pgbadger

do-install: build
		@${DO_NADA}

.include <bsd.port.mk>
