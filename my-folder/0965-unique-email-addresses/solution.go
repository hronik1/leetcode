func numUniqueEmails(emails []string) int {
    normalizedEmails := map[string]bool{}
    for _, email := range emails {
        normalizedEmail := normalizeEmail(email)
        normalizedEmails[normalizedEmail] = true
    }
    
    return len(normalizedEmails)
}

func normalizeEmail(email string) string {
    parts := strings.Split(email, "@")
    // ensure two parts
    prefix := strings.Split(parts[0], "+")
    normalizedPrefix := strings.Join(strings.Split(prefix[0], "."), "")
    return fmt.Sprintf("%s@%s", normalizedPrefix, parts[1])
}
