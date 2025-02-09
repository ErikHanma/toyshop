#!/bin/bash

exclusions=("$@")  # Получаем список исключений из аргументов

print_directory() {
    local dir="$1"
    local indent="$2"

    local dir_name=$(basename "$dir")
    # Проверяем, исключена ли текущая директория
    for excl in "${exclusions[@]}"; do
        if [ "$excl" == "$dir_name" ]; then
            return
        fi
    done

    echo "${indent}${dir_name}/"
    
    for file in "$dir"/*; do
        local filename=$(basename "$file")
        local skip=0

        # Проверяем, исключён ли текущий файл или директория
        for excl in "${exclusions[@]}"; do
            if [ "$excl" == "$filename" ]; then
                skip=1
                break
            fi
        done
        if [ $skip -eq 1 ]; then
            continue
        fi

        if [ -d "$file" ]; then
            print_directory "$file" "  ${indent}"
        elif [ -f "$file" ]; then
            echo "${indent}  ${filename}"
            echo "${indent}  --------------------"
            cat "$file"
            echo ""
            echo "${indent}  --------------------"
        fi
    done
}

print_directory "." ""