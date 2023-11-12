IMAGE_NAME = go-fiber-mongodb
ARTIFACTS_FOLDER = artifacts
TAR_FILE = $(ARTIFACTS_FOLDER)/$(IMAGE_NAME).tar

save-image:
	docker save -o $(TAR_FILE) $(IMAGE_NAME)


