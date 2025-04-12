import secrets
import string
from pathlib import Path

def generate_secure_password(length=24):
    """Generate a secure random password."""
    alphabet = string.ascii_letters + string.digits
    return ''.join(secrets.choice(alphabet) for _ in range(length))

def update_env_file(env_path):
    """Update .env file with generated credentials if they are not set."""
    # Read existing .env file
    env_vars = {}
    if env_path.exists():
        with open(env_path, 'r') as f:
            for line in f:
                line = line.strip()
                if line and not line.startswith('#'):
                    if '=' in line:
                        key, value = line.split('=', 1)
                        env_vars[key.strip()] = value.strip()

    # Define required variables and their default values
    required_vars = {
        'REDIS_PASSWORD': generate_secure_password(),
        'REDIS_COMMANDER_PASSWORD': generate_secure_password(),
        'REDIS_COMMANDER_USER': 'admin',
        'REDIS_HOST': 'redis-service',
        'REDIS_PORT': '6379',
        'REDIS_DB': '0',
        'REDIS_COMMANDER_PORT': '8081',
        'SORAME_SERVICE_PORT': '3000',
        'NETWORK_NAME': 'sorame-network'
    }

    # Update env_vars with missing variables
    updated = False
    for key, default_value in required_vars.items():
        if key not in env_vars or not env_vars[key]:
            env_vars[key] = default_value
            updated = True
            print(f"Generated {key}: {default_value}")

    # Write back to .env file if updates were made
    if updated:
        with open(env_path, 'w') as f:
            # Write Redis Configuration
            f.write("# Redis Configuration\n")
            f.write(f"REDIS_HOST={env_vars['REDIS_HOST']}\n")
            f.write(f"REDIS_PORT={env_vars['REDIS_PORT']}\n")
            f.write(f"REDIS_PASSWORD={env_vars['REDIS_PASSWORD']}\n\n")
            f.write(f"REDIS_DB={env_vars['REDIS_DB']}\n\n")
            
            # Write Redis Commander Configuration
            f.write("# Redis Commander Configuration\n")
            f.write(f"REDIS_COMMANDER_HOST=redis-commander\n")
            f.write(f"REDIS_COMMANDER_PORT={env_vars['REDIS_COMMANDER_PORT']}\n")
            f.write(f"REDIS_COMMANDER_USER={env_vars['REDIS_COMMANDER_USER']}\n")
            f.write(f"REDIS_COMMANDER_PASSWORD={env_vars['REDIS_COMMANDER_PASSWORD']}\n\n")
            
            # Write Sorame Service Configuration
            f.write("# Sorame Service Configuration\n")
            f.write(f"SORAME_SERVICE_PORT={env_vars['SORAME_SERVICE_PORT']}\n\n")
            
            # Write Network Configuration
            f.write("# Network Configuration\n")
            f.write(f"NETWORK_NAME={env_vars['NETWORK_NAME']}\n")
        
        print("\n.env file has been updated with new credentials.")
    else:
        print("\nAll credentials are already set. No changes were made.")

if __name__ == "__main__":
    env_path = Path('.env')
    update_env_file(env_path) 