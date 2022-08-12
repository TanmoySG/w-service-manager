# For more information, please refer to https://aka.ms/vscode-docker-python
FROM kuralabs/python3-dev

EXPOSE 5002

# Keeps Python from generating .pyc files in the container
ENV PYTHONDONTWRITEBYTECODE=1

# Turns off buffering for easier container logging
ENV PYTHONUNBUFFERED=1

# Install pip requirements
COPY requirements.txt .
RUN python3 -m pip install -r requirements.txt

WORKDIR /app
COPY . /app

RUN mkdir configurations

# ENTRYPOINT ["python3"]
# CMD ["app.py"]

CMD [ "sh", "./entrypoint.sh" ]
