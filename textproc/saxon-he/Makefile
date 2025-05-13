PORTNAME=	saxon-he
DISTVERSION=	12-6
CATEGORIES=	textproc java
MASTER_SITES=	https://github.com/${GH_ACCOUNT}/${GH_PROJECT}/releases/download/SaxonHE${DISTVERSION}/
DISTNAME=	SaxonHE${DISTVERSION}J

MAINTAINER=	thierry@FreeBSD.org
COMMENT=	XSLT 3.0 / XPath 3.1 / XQuery 3.1 processor for Java
WWW=		https://www.saxonica.com/

LICENSE=	MPL20
LICENSE_FILE=	${WRKSRC}/notices/LICENSE.txt

#USE_GITHUB=	yes
GH_ACCOUNT=	Saxonica
GH_PROJECT=	Saxon-HE

USES=		java zip

NO_WRKSUBDIR=	yes
NO_ARCH=	yes
NO_BUILD=	yes

JARFILES=	${JARFILES1} ${JARFILES2}
JARFILES1=	saxon-he-${PORTVERSION}.jar \
		saxon-he-test-${PORTVERSION}.jar \
		saxon-he-xqj-${PORTVERSION}.jar
JARFILES2=	jline-2.14.6.jar \
		xmlresolver-5.3.1.jar \
		xmlresolver-5.3.1-data.jar
PLIST_FILES=	${JARFILES:S,^,%%JAVAJARDIR%%/,} ${SCRIPTFILES}

SUB_FILES=	saxon.sh
SUB_LIST=	SAXON_JARS="${JARFILES:S,^,${JAVAJARDIR}/,}"
SCRIPTFILES=	bin/saxon bin/saxon-xquery

do-install:
	@${MKDIR} ${STAGEDIR}${JAVAJARDIR}
	${INSTALL_DATA} ${JARFILES1:S,^,${WRKSRC}/,} ${STAGEDIR}${JAVAJARDIR}
	${INSTALL_DATA} ${JARFILES2:S,^,${WRKSRC}/lib/,} ${STAGEDIR}${JAVAJARDIR}
.for scriptfile in ${SCRIPTFILES}
	${INSTALL_SCRIPT} ${WRKDIR}/saxon.sh ${STAGEDIR}${PREFIX}/${scriptfile}
.endfor

.include <bsd.port.mk>
