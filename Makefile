.PHONY: flatpak-build flatpak-install flatpak-run flatpak-clean flatpak-bundle icon clean help check-deps install-deps

# Colors for output
RED = \033[0;31m
GREEN = \033[0;32m
YELLOW = \033[1;33m
BLUE = \033[0;34m
NC = \033[0m # No Color

# Default target
help:
	@echo "BONK Flatpak Build Targets:"
	@echo "  check-deps    - Check if required dependencies are installed"
	@echo "  install-deps  - Install required Flatpak runtime and SDK"
	@echo "  icon          - Create application icon"
	@echo "  flatpak-build - Build Flatpak package"
	@echo "  flatpak-install - Install Flatpak locally"
	@echo "  flatpak-run   - Run the installed Flatpak"
	@echo "  flatpak-bundle - Create distributable .flatpak file"
	@echo "  flatpak-clean - Clean build artifacts"
	@echo "  clean         - Clean all generated files"
	@echo "  all           - Build, install and test everything"
	@echo "  test          - Test the installation"

# Application details
APP_ID = net.jorjai.BONK
MANIFEST = $(APP_ID).yml
BUILD_DIR = build-dir
REPO_DIR = flatpak-repo

# Check dependencies
check-deps:
	@echo -e "$(BLUE)Checking dependencies...$(NC)"
	@command -v flatpak-builder >/dev/null 2>&1 || { \
		echo -e "$(RED)Error: flatpak-builder is not installed.$(NC)"; \
		echo "Install it with:"; \
		echo "  Ubuntu/Debian: sudo apt install flatpak-builder"; \
		echo "  Fedora: sudo dnf install flatpak-builder"; \
		echo "  Arch: sudo pacman -S flatpak-builder"; \
		exit 1; \
	}
	@flatpak remotes | grep -q flathub || { \
		echo -e "$(YELLOW)Adding Flathub remote...$(NC)"; \
		flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo; \
	}
	@echo -e "$(GREEN)✅ Dependencies check passed!$(NC)"

# Install required runtime and SDK
install-deps: check-deps
	@echo -e "$(BLUE)Installing runtime and SDK...$(NC)"
	@flatpak install -y flathub org.freedesktop.Platform//24.08 2>/dev/null || true
	@flatpak install -y flathub org.freedesktop.Sdk//24.08 2>/dev/null || true
	@flatpak install -y flathub org.freedesktop.Sdk.Extension.golang//24.08 2>/dev/null || true
	@echo -e "$(GREEN)✅ Runtime and SDK installed!$(NC)"

# Create icon
icon:
	@echo -e "$(BLUE)Creating application icon...$(NC)"
	@./create-icon.sh
	@echo -e "$(GREEN)✅ Icon created!$(NC)"

# Build the Flatpak
flatpak-build: install-deps
	@echo -e "$(BLUE)Building Flatpak package...$(NC)"
	@flatpak-builder --force-clean $(BUILD_DIR) $(MANIFEST)
	@echo -e "$(BLUE)Creating Flatpak repository...$(NC)"
	@flatpak-builder --repo=$(REPO_DIR) --force-clean $(BUILD_DIR) $(MANIFEST)
	@echo -e "$(GREEN)✅ Flatpak built successfully!$(NC)"

# Install locally
flatpak-install: flatpak-build
	@echo -e "$(BLUE)Installing application locally...$(NC)"
	@flatpak --user remote-add --no-gpg-verify --if-not-exists bonk-repo $(REPO_DIR) 2>/dev/null || true
	@flatpak --user install -y bonk-repo $(APP_ID)
	@echo -e "$(GREEN)✅ BONK Flatpak installed successfully!$(NC)"

# Run the application
flatpak-run:
	@echo -e "$(BLUE)Running BONK Flatpak...$(NC)"
	@flatpak run $(APP_ID)

# Create bundle for distribution
flatpak-bundle: flatpak-build
	@echo -e "$(BLUE)Creating distributable bundle...$(NC)"
	@flatpak build-bundle $(REPO_DIR) bonk.flatpak $(APP_ID)
	@echo -e "$(GREEN)✅ Bundle created: bonk.flatpak$(NC)"

# Clean Flatpak build artifacts
flatpak-clean:
	@echo -e "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -rf $(BUILD_DIR) $(REPO_DIR) .flatpak-builder

# Clean all generated files
clean: flatpak-clean
	@echo -e "$(YELLOW)Cleaning all generated files...$(NC)"
	@rm -f bonk.flatpak icon.png icon_placeholder.txt

# Quick build and install with icon
all: icon flatpak-install test

# Test the installation
test: 
	@echo -e "$(BLUE)Testing BONK installation...$(NC)"
	@flatpak run $(APP_ID) --version
	@echo -e "$(GREEN)✅ BONK Flatpak is working!$(NC)"