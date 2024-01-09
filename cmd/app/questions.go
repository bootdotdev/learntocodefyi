package main

type QuestionType string

const (
	ChoiceQuestionType  QuestionType = "choice"
	CountryQuestionType QuestionType = "country"
)

type Question struct {
	ID              string
	QuestionType    QuestionType
	ChoiceQuestion  *ChoiceQuestion
	CountryQuestion *CountryQuestion
}

type ChoiceQuestion struct {
	Question     string
	Answers      []string
	MaxSelection int
}

type CountryQuestion struct {
	Question string
}

func getQuestions() []Question {
	return []Question{
		{
			ID:           "f7a0817c-899e-4452-ba19-65846611ad2e",
			QuestionType: CountryQuestionType,
			CountryQuestion: &CountryQuestion{
				Question: "Where do you live?",
			},
		},
		{
			ID:           "8a585909-c337-4e53-a098-1bfbd84f0696",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How old are you?",
				Answers: []string{
					"Under 18",
					"18-24",
					"25-34",
					"35-44",
					"45-54",
					"Over 54",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "b3da3188-966c-42f1-b264-447177038be1",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How much formal education have you received? (Doesn't need to be programming related)",
				Answers: []string{
					"Primary/elementary school",
					"Secondary school (High school)",
					"Some college/university study without earning a degree",
					"Associate degree",
					"Bachelor's degree",
					"Master's degree",
					"Doctoral degree",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "a918ce09-4237-4cbb-bb63-1d7af8544c53",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How many years of paid professional programming experience do you have, if any?",
				Answers: []string{
					"0",
					"<1",
					"1-2",
					"2-3",
					"3-4",
					"4-5",
					"5-6",
					"6-7",
					"7-8",
					"8+",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "7d364d99-01b1-4188-8558-4d03018ffeff",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Have you completed any of the following types of certified informal education?",
				Answers: []string{
					"Physical Bootcamp",
					"Online Bootcamp (With live instruction)",
					"Online Bootcamp (Fully self-paced, no live instruction)",
					"University certificate program",
					"Online course",
				},
				MaxSelection: -1,
			},
		},
		{
			ID:           "08fd2ac4-d3fd-481e-81c6-fca6f4ebfa97",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How long did you spend learning to code before you landed your first programming job? Only count months where you spent at least 10 hours per week learning. If you don't have a programming job, how long have you been studying in total?",
				Answers: []string{
					"<3 months",
					"3-6 months",
					"6-9 months",
					"9-12 months",
					"12-18 months",
					"18-24 months",
					"2-3 years",
					"3-4 years",
					"4-6 years",
					"6+ years",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "4e609528-e46a-46e8-bb7e-4fae878aaa17",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which do you feel have been the most effective resources for learning a new programming language or framework?",
				Answers: []string{
					"Official documentation",
					"Books",
					"YouTube videos",
					"Coding live streams (Twitch, etc.)",
					"Podcasts",
					"Blog posts and articles",
					"Video courses",
					"Interactive courses",
					"Text-based courses",
					"Coding challenges",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "2f2ef4f7-4b10-4271-88c2-fab66a89e090",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which do you feel have been the most effective resources for learning a new tool or library?",
				Answers: []string{
					"Official documentation",
					"Books",
					"YouTube videos",
					"Coding live streams (Twitch, etc.)",
					"Podcasts",
					"Blog posts and articles",
					"Video courses",
					"Interactive courses",
					"Text-based courses",
					"Coding challenges",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "11b6b646-424b-45fa-8167-3c39fc4df105",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Where do you get most of your programming news?",
				Answers: []string{
					"Twitter/X",
					"Reddit",
					"Hacker News",
					"YouTube",
					"Podcasts",
					"Newsletters",
					"Personal Blogs",
					"Blogging platforms",
					"LinkedIn",
					"Facebook",
					"Instagram",
					"Snapchat",
					"TikTok",
					"Other",
				},
				MaxSelection: 5,
			},
		},
		{
			ID:           "2e441a8c-6b07-492e-8cbc-0161213310c5",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How much do you (or did you) hope to earn in your first full time programming job?",
				Answers: []string{
					"<$20,000 / year",
					"$20,000 - $40,000 / year",
					"$40,000 - $60,000 / year",
					"$60,000 - $80,000 / year",
					"$80,000 - $100,000 / year",
					">$100,000 / year",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "794c70a1-b66a-46e1-acdf-daeda0d4cb5b",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which types of developer roles are you most interested in?",
				Answers: []string{
					"Front-end web",
					"Back-end web",
					"Full-stack web",
					"Mobile",
					"Desktop or enterprise applications",
					"DevOps / SysAdmin / SRE / Cloud Infrastructure",
					"Data scientist or machine learning specialist",
					"Data engineer",
					"Blockchain",
					"Database Administrator",
					"Embedded applications or devices",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "05c6cafc-0b4d-4da9-9715-76870873243c",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which programming languages are you most interested in working with professionally?",
				Answers: []string{
					"JavaScript",
					"HTML/CSS",
					"Python",
					"SQL",
					"TypeScript",
					"Java",
					"C#",
					"C++",
					"C",
					"PHP",
					"Go",
					"Rust",
					"Ruby",
					"Kotlin",
					"Dart",
					"Lua",
					"Swift",
					"Scala",
					"OCaml",
					"Zig",
					"other",
				},
				MaxSelection: 5,
			},
		},
		{
			ID:           "05d7ec34-479b-4769-bc4d-ed5a59113b3d",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which programming languages are you most interested in learning personally?",
				Answers: []string{
					"JavaScript",
					"HTML/CSS",
					"Python",
					"SQL",
					"TypeScript",
					"Java",
					"C#",
					"C++",
					"C",
					"PHP",
					"Go",
					"Rust",
					"Ruby",
					"Kotlin",
					"Dart",
					"Lua",
					"Swift",
					"Scala",
					"OCaml",
					"Zig",
					"other",
				},
				MaxSelection: 5,
			},
		},
		{
			ID:           "63626531-f76b-448b-834c-81344afe7d73",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How do you balance your time in regard to passive and active learning? (Passive learning is watching videos, reading books, etc. Active learning is building projects, coding challenges, etc.)",
				Answers: []string{
					"10% Passive, 90% Active",
					"25% Passive, 75% Active",
					"50% Passive, 50% Active",
					"75% Passive, 25% Active",
					"90% Passive, 10% Active",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "1845f5ff-dabb-4ff7-b0b8-88bd22fd17b1",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which of the following do you use most for help when you're stuck?",
				Answers: []string{
					"Official documentation",
					"Web Search (e.g. Google)",
					"AI Chat (e.g ChatGPT)",
					"Forum (e.g. Stack Overflow, Reddit)",
					"Video Search (e.g. YouTube)",
					"Live community (e.g. Discord, Slack)",
					"A personal friend, family member, or colleague",
				},
				MaxSelection: 2,
			},
		},
		{
			ID:           "2327236d-32c2-494c-8546-763585f1cbbe",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which of the following have been most effective for networking with professional developers?",
				Answers: []string{
					"Local Meetups",
					"Online Meetups",
					"Conferences",
					"Slack communities",
					"Discord communities",
					"Social Media",
				},
				MaxSelection: 2,
			},
		},
		{
			ID:           "0881f3f0-9e43-4455-9480-c2ed2ff3c2f1",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How much money have you spent on all educational resources out of your own pocket?",
				Answers: []string{
					"$0",
					"$1-$500",
					"$500-$1,999",
					"$2,000-$9,999",
					"$10,000-$49,999",
					"$50,000+",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "ea0c0dcb-f901-45bd-a4ec-84c10feaf78f",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How much money has your employer spent on educational resources for you?",
				Answers: []string{
					"$0",
					"$1-$500",
					"$500-$1,999",
					"$2,000-$9,999",
					"$10,000-$49,999",
					"$50,000+",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "2eec14d9-ca48-40aa-93ad-d6d3aa4d78a7",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which do you find to be most distracting when you're trying to focus on code?",
				Answers: []string{
					"Social Videos (YouTube, TikTok, etc.)",
					"Live Streams (Twitch, etc.)",
					"Social Media (Twitter, Facebook, etc.)",
					"News (Reddit, Hacker News, etc.)",
					"Live Chat (Discord, Slack, etc.)",
					"Home Life (Family, Roommates, etc.)",
					"Video Games",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "01f659b6-3c9f-4e75-b3f5-e012d4800bfc",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "What do you find most difficult about learning to code?",
				Answers: []string{
					"Finding the time",
					"Losing interest or getting bored",
					"Getting help with specific problems when I'm stuck",
					"Knowing what to learn next",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "b8531043-0a1e-4090-91d6-5b213c5e6304",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which operating system do you primarily code on?",
				Answers: []string{
					"Windows",
					"Windows (WSL)",
					"MacOS",
					"Linux",
					"ChromeOS",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "74bfbfe0-0ddf-4363-bcb0-3b977eb093cc",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which do you prefer to code in?",
				Answers: []string{
					"Visual Studio Code",
					"Visual Studio",
					"IntelliJ IDEA",
					"NotePad++",
					"Vim",
					"PyCharm",
					"Jupyter Notebooks",
					"Sublime Text",
					"Neovim",
					"Eclipse",
					"Xcode",
					"Android Studio",
					"Emacs",
					"Goland",
					"Other",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "1a74c95b-5bea-4a21-ba1d-d4c3fc35d2aa",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which is the hardest part of building large projects from scratch?",
				Answers: []string{
					"I can follow guided tutorials, but don't know where to start on my own",
					"I don't know how to put basic concepts together to build a full project",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "b20c615c-2455-4ded-b05e-9183cc891f8a",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which are you most worried about?",
				Answers: []string{
					"I'm worried AI will replace programmers",
					"I'm worried I'm not smart enough to be a programmer",
					"I'm worried I'll be discriminated against",
					"I'm worried I don't have enough time in my day to study",
					"I'm worried it will take too long to learn to code",
					"I'm worried that it's too expensive to learn to code",
					"I'm worried I'm not in a good location for coding jobs",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "f05fc0bb-4da7-4ea1-97c9-83d34a62b59f",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "What worriess you most about a new learning resource?",
				Answers: []string{
					"I'm worried it won't be hands-on enough",
					"I'm worried it will be too hard or too easy",
					"I'm worried it will be too expensive",
					"I'm worried it will take too long",
					"I'm worried it will be too boring and I'll lose interest",
					"I'm worried it will have innacurate or misleading information",
					"I'm worried it won't teach clearly enough",
					"I'm worried it won't go in depth enough",
				},
				MaxSelection: 3,
			},
		},
		{
			ID:           "f1c91a48-1977-4281-9cba-e8fb3745b671",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which do you value most in a programming job?",
				Answers: []string{
					"High salary",
					"Health insurance",
					"Remote work",
					"Flexible hours",
					"Equity in the company",
					"Fun environment (ping pong, free food, etc.)",
				},
				MaxSelection: 2,
			},
		},
		{
			ID:           "be5f094c-3664-4560-8d80-4df432cccf1a",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "For your first job, which are you (or were you) most interested in?",
				Answers: []string{
					"Fully remote, same country",
					"Fully remote, different country",
					"Hybrid",
					"In office",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "3adcddbe-feb5-4eb8-bcba-39cf4be629eb",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "For your first job, do you plan on applying (or did you) to in-person jobs as well?",
				Answers: []string{
					"Yes",
					"No",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "a6cce526-baad-48ab-abe8-da432b932f74",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Which are you most worried about when it comes to qualifying for a programming job?",
				Answers: []string{
					"I'm worried I don't know how to build practical applications from scratch",
					"I'm worried I don't understand theory and fundamentals well enough",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "6d457ed9-5607-4cfb-9302-ccc4d421f5c2",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Are you actively looking for a programming job?",
				Answers: []string{
					"Yes, seriously",
					"Yes, casually",
					"No",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "eac750f5-5039-4dfd-914b-f20bde9b0290",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "How many programming jobs do you apply to per week while you're actively looking?",
				Answers: []string{
					"0",
					"1-5",
					"5-10",
					"10-20",
					"20-40",
					"40-80",
					"80-160",
					"160+",
				},
				MaxSelection: 1,
			},
		},
		{
			ID:           "70f59f95-9181-46c3-8a0e-03abd75d72a9",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "Where are your favorite places to look for programming jobs?",
				Answers: []string{
					"Big job boards (LinkedIn, Indeed, etc.)",
					"Small niche job boards",
					"Slack/Discord communities",
					"Personal network",
					"Physical Meetups",
					"Job fairs",
					"Conferences",
					"Cold outreach (dms, emails, etc.)",
					"Recruiters",
				},
				MaxSelection: 3,
			},
		},

		{
			ID:           "0893ca93-9be5-4411-9c7f-0f1299190911",
			QuestionType: ChoiceQuestionType,
			ChoiceQuestion: &ChoiceQuestion{
				Question: "What are you (or were you) most worried about with applying to for your first programming job?",
				Answers: []string{
					"I don't have the practical skills",
					"I don't have the theoretical knowledge",
					"I don't have good portfolio projects",
					"I don't have a formal degree",
					"I don't have any professional experience",
					"I don't have a network",
					"I don't have a good personal brand or social media presence",
					"I don't know the right languages, frameworks, or tools",
				},
				MaxSelection: 3,
			},
		},
	}
}
