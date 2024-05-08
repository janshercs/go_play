import csv
import os


def process_orders():
    TRANSFER_TO_SOURCE = set()
    # Opening JSON file
    with open('./transfer_to_source.csv') as f:
        for row in csv.reader(f):
            TRANSFER_TO_SOURCE.add(row[0])

    PROCESSED = [
        ['reference', 'transfer_to_source', 'status',
            'payment_method_code', 'updated_at']
    ]

    no_invoice_file = 'payments_with_SAP_invoice_status_2024-05-05_to_2025-05-06_without_invoice_row.csv'
    with open(no_invoice_file, 'r') as f:
        for row in csv.reader(f):
            PROCESSED.append(
                [row[0], str(row[0] in TRANSFER_TO_SOURCE),
                 row[1], row[2], row[3]]
            )

    PROCESSED.sort(key=lambda x: x[4])

    with open('output.csv', 'w', newline='') as f:
        csv_writer = csv.writer(f)
        # Write rows to the CSV file
        for row in PROCESSED:
            csv_writer.writerow(row)


def split_csv(input_file, output_folder, chunk_sizes=[50, 1000, 2000, 3000, 4000, 5000]):
    # Create output folder if it doesn't exist
    os.makedirs(output_folder, exist_ok=True)

    with open(input_file, 'r') as f:
        i = 0
        header = f.readline()  # Read the header
        chunk_number = 1
        chunk_size = chunk_sizes[i]
        current_chunk_size = 0
        current_chunk_rows = [header]

        for line in f:
            current_chunk_rows.append(line)
            current_chunk_size += 1

            # If the current chunk size reaches the specified chunk size, write it to a new file
            if current_chunk_size == chunk_size:
                with open(os.path.join(output_folder, f'chunk_{chunk_number}_{chunk_size}_records.csv'), 'w') as chunk_file:
                    chunk_file.writelines(current_chunk_rows)

                # Reset variables for the next chunk
                current_chunk_rows = [header]
                current_chunk_size = 0
                chunk_number += 1
                i += 1
                chunk_size = chunk_sizes[i]

        # Write the remaining rows to the last chunk file
        if current_chunk_rows:
            with open(os.path.join(output_folder, f'chunk_{i}_{chunk_number}.csv'), 'w') as chunk_file:
                chunk_file.writelines(current_chunk_rows)


# Example usage

input_file = 'output.csv'
output_folder = 'output'
chunk_sizes = [50, 1000, 2000, 3000, 4000, 5000,
               5000, 5000, 5000, 5000, 5000]
split_csv(input_file, output_folder, chunk_sizes)

# process_orders()
