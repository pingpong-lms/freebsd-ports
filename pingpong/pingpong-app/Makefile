# Created by: girgen
#
PORTNAME=	pingpong-app
PORTVERSION=	0.28
CATEGORIES=	pingpong
MASTER_SITES=	#
DISTFILES=	#
EXTRACT_ONLY=	#

MAINTAINER=	girgen@pingpong.net
COMMENT=	The meta-port for PING PONG application packages

NO_BUILD=	yes
USE_JAVA=	yes
USE_RUBY=	yes
JAVA_VERSION=	1.8+

PLIST_FILES=	sbin/ppctl \
		share/ffmpeg/libx264-slow.ffpreset

USE_APACHE_RUN=	24
USES=		perl5 python:3.7+

RUN_DEPENDS= \
	jsvc:devel/apache-commons-daemon \
	convert:graphics/ImageMagick7-nox11 \
	p5-Pg>0:databases/p5-Pg \
	p5-XML-Writer>0:textproc/p5-XML-Writer \
	${LOCALBASE}/${APACHEMODDIR}/mod_jk.so:www/mod_jk \
	${LOCALBASE}/${APACHEMODDIR}/mod_xsendfile.so:www/mod_xsendfile \
	latex2html:textproc/latex2html \
	html2ps:print/html2ps \
	ant:devel/apache-ant \
	npm:www/npm \
	clamd:security/clamav \
	ffmpeg:multimedia/ffmpeg \
	webalizer:www/webalizer \
	shibd:security/shibboleth-sp \
	saxon:textproc/saxon-devel \
	jq:textproc/jq \
	xml:textproc/xmlstarlet \
	xmlsec1:security/xmlsec1 \
	${LOCALBASE}/share/fonts/webfonts/arial.ttf:x11-fonts/webfonts \
	wkhtmltopdf:converters/wkhtmltopdf \
	flask:www/py-flask \
	${PYTHON_PKGNAMEPREFIX}dateutil>0:devel/py-dateutil@${PY_FLAVOR} \
	${PYTHON_PKGNAMEPREFIX}requests>=2.20.0:www/py-requests@${PY_FLAVOR} \
	automake:devel/automake \
	${LOCALBASE}/share/pingpong-generic/installed.stamp:pingpong/pingpong-generic \
	${LOCALBASE}/share/pingpong-aspell/installed.stamp:pingpong/pingpong-aspell

do-install:
	${LN} -sf ../../../opt/pingpong/libexec/ppctl ${STAGEDIR}${PREFIX}/sbin/ppctl
	${MKDIR} ${STAGEDIR}${PREFIX}/share/ffmpeg && \
	${INSTALL_DATA} ${FILESDIR}/libx264-slow.ffpreset ${STAGEDIR}${PREFIX}/share/ffmpeg

.include <bsd.port.mk>
