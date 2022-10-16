import { generateNanoId } from "../index";

test('Generate Nano ID', () => {
    expect(generateNanoId(12)).toBeDefined()
})