FROM loads/alpine:3.8

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR                 /app
COPY ./manifest/config/config.prod.yaml  $WORKDIR/config.yaml
ADD ./main $WORKDIR
RUN chmod +x $WORKDIR

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
#CMD GF_GCFG_FILE=config.prod.toml
CMD GF_GCFG_FILE=config.prod.toml; ./main
