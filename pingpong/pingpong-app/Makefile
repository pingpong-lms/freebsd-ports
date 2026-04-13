# Created by: girgen
#
PORTNAME=	pingpong-app
PORTVERSION=	2.4
CATEGORIES=	pingpong
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for PING PONG application packages

NO_BUILD=	yes
JAVA_VERSION=	25
JAVA_FLAVOR=	headless

PLIST_FILES=	sbin/ppctl \
		share/ffmpeg/libx264-slow.ffpreset

USES=		perl5 python ruby:run magick:7,nox11,run apache:2.4,run java

RUN_DEPENDS= \
	${LOCALBASE}/openjdk8/bin/java:java/openjdk8 \
	${LOCALBASE}/openjdk21/bin/java:java/openjdk21 \
	jsvc:devel/apache-commons-daemon \
	p5-Pg>0:databases/p5-Pg \
	p5-XML-Writer>0:textproc/p5-XML-Writer \
	${LOCALBASE}/${APACHEMODDIR}/mod_jk.so:www/mod_jk \
	${LOCALBASE}/${APACHEMODDIR}/mod_xsendfile.so:www/mod_xsendfile \
	latex2html:textproc/latex2html \
	ant:devel/apache-ant \
	npm-node18>=0:www/npm-node18 \
	clamd:security/clamav \
	ffmpeg:multimedia/ffmpeg \
	webalizer:www/webalizer \
	shibd:security/shibboleth-sp \
	saxon:textproc/saxon-he \
	jq:textproc/jq \
	xml:textproc/xmlstarlet \
	xmlsec1:security/xmlsec1 \
	${LOCALBASE}/share/fonts/webfonts/arial.ttf:x11-fonts/webfonts \
	wkhtmltopdf:converters/wkhtmltopdf \
	weasyprint:print/py-weasyprint \
	certbot:security/py-certbot \
	flask:www/py-flask \
	filebeat:sysutils/beats8 \
	${PYTHON_PKGNAMEPREFIX}python-dateutil>0:devel/py-python-dateutil@${PY_FLAVOR} \
	${PYTHON_PKGNAMEPREFIX}requests>=2.20.0:www/py-requests@${PY_FLAVOR} \
	automake:devel/automake \
	${LOCALBASE}/share/pingpong-generic/installed.stamp:pingpong/pingpong-generic \
	${LOCALBASE}/share/pingpong-aspell/installed.stamp:pingpong/pingpong-aspell

do-install:
	${LN} -sf ../../../opt/pingpong/libexec/ppctl ${STAGEDIR}${PREFIX}/sbin/ppctl
	${MKDIR} ${STAGEDIR}${PREFIX}/share/ffmpeg && \
	${INSTALL_DATA} ${FILESDIR}/libx264-slow.ffpreset ${STAGEDIR}${PREFIX}/share/ffmpeg

.include <bsd.port.mk>
