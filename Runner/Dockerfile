FROM python:3.7.4-alpine3.10
RUN pip install iso4217parse
ADD runner.py /
EXPOSE 3000
ENTRYPOINT [ "python", "./runner.py" ]
